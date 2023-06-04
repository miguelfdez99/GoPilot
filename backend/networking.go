package backend

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type NetworkInterface struct {
	Name   string
	IP     string
	Status string
}

func (b *Backend) GetNetworkInterfaces() ([]NetworkInterface, error) {
	cmd := exec.Command("ip", "-br", "address", "show")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error running ip command: %w", err)
	}

	lines := strings.Split(out.String(), "\n")
	interfaces := make([]NetworkInterface, 0, len(lines))

	ipRegex := regexp.MustCompile(`(?m)\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Fields(line)
			iface := NetworkInterface{
				Name:   fields[0],
				Status: fields[1],
			}

			cmd := exec.Command("ip", "-br", "address", "show", iface.Name)
			var out bytes.Buffer
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				return nil, fmt.Errorf("error running ip command: %w", err)
			}

			ipMatches := ipRegex.FindStringSubmatch(out.String())
			if len(ipMatches) > 0 {
				iface.IP = ipMatches[0]
			} else {
				iface.IP = "N/A"
			}

			interfaces = append(interfaces, iface)
		}
	}
	return interfaces, nil
}

func (b *Backend) SetInterfaceStatus(name string, status string) error {
	if status != "up" && status != "down" {
		return fmt.Errorf("invalid status: %s", status)
	}

	cmd := exec.Command("ip", "link", "set", name, status)
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing command: %v, stderr: %s", err, errBuf.String())
	}

	return nil
}
