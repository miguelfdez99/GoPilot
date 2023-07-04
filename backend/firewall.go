package backend

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var distributionsMap = map[string]func(string) error{
	"ubuntu": addUfwRule,
	"debian": addUfwRule,
}

type TrafficData struct {
	Netid       string `json:"netid"`
	LocalAddr   string `json:"localAddr"`
	PeerAddr    string `json:"peerAddr"`
	Application string `json:"application"`
}

func isValidRule(rule string) bool {
	if rule == "" {
		return false
	}

	forbiddenChars := []string{";", "&", "|", "`", "$"}
	for _, char := range forbiddenChars {
		if strings.Contains(rule, char) {
			return false
		}
	}

	return true
}

func (b *Backend) ListUfw() ([]string, error) {
	cmd := exec.Command("ufw", "status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	if lines[0] == "" {
		return lines[1:], nil
	}
	return lines, nil
}

func addUfwRule(rule string) error {
	args := append([]string{"ufw"}, strings.Fields(rule)...)
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeUfwRule(rule string) error {
	args := append([]string{"ufw", "delete"}, strings.Fields(rule)...)
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (b *Backend) GetFirewallStatus() (string, error) {
	cmd := exec.Command("ufw", "status")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	status := strings.Split(string(output), "\n")[0]
	fmt.Println(status)
	return status, nil
}

func (b *Backend) SetFirewallStatus(status string) error {
	if status != "active" && status != "inactive" {
		return errors.New("invalid status")
	}

	var cmd *exec.Cmd
	if status == "active" {
		cmd = exec.Command("ufw", "enable")
	} else {
		cmd = exec.Command("ufw", "disable")
	}

	b.logger.Info(fmt.Sprintf("Firewall %s", status))

	return cmd.Run()
}

func (b *Backend) GetDefaultPolicy(direction string) (string, error) {
	if direction != "incoming" && direction != "outgoing" {
		return "", errors.New("invalid direction")
	}

	cmd := exec.Command("ufw", "default", direction)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	policy := strings.Split(string(output), " ")[4]
	return policy, nil
}

func (b *Backend) SetDefaultPolicy(direction, policy string) error {
	if direction != "incoming" && direction != "outgoing" {
		return errors.New("invalid direction")
	}
	if policy != "allow" && policy != "deny" {
		return errors.New("invalid policy")
	}

	cmd := exec.Command("ufw", "default", policy, direction)
	return cmd.Run()
}

func (b *Backend) FetchTrafficData() ([]TrafficData, error) {
	cmd := exec.Command("ss", "-tunlp")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(output))

	lines := strings.Split(string(output), "\n")

	trafficData := make([]TrafficData, 0, len(lines))

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		netid := fields[0]
		localAddr := fields[3]
		peerAddr := fields[4]

		processInfo := fields[5]
		processSplit := strings.Split(processInfo, "\"")
		application := ""
		if len(processSplit) > 1 {
			application = processSplit[1]
		}

		if netid == "" || localAddr == "" || peerAddr == "" {
			continue
		}

		trafficData = append(trafficData, TrafficData{
			Netid:       netid,
			LocalAddr:   localAddr,
			PeerAddr:    peerAddr,
			Application: application,
		})
	}

	return trafficData, nil
}
