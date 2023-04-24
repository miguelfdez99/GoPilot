package backend

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func (b *Backend) CheckAdmin() bool {
	uid := os.Getuid()
	if uid != 0 {
		fmt.Println("You are not an admin")
		return false
	} else {
		fmt.Println("You are an admin")
		return true
	}
}

func ExtractFirstParams(input string) []string {
	distribution, err := getLinuxDistribution()
	if err != nil {
		return nil
	}

	lines := strings.Split(input, "\n")
	params := make([]string, len(lines))
	for i, line := range lines {
		if distribution == "debian" || distribution == "ubuntu" {
			fields := strings.Split(line, "/")
			params[i] = fields[0]
		} else if distribution == "fedora" || distribution == "centos" || distribution == "rhel" {
			fields := strings.Split(line, ".")
			params[i] = fields[0]
		} else if distribution == "arch" {
			fields := strings.Split(line, " ")
			params[i] = fields[0]
		}
	}
	return params
}

func getLinuxDistribution() (string, error) {
	fileContent, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(fileContent))
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID_LIKE" {
			return strings.Trim(parts[1], "\""), nil
		}
	}

	// If ID_LIKE is not present, try ID(Fedora does not have ID_LIKE)
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID" {
			return strings.Trim(parts[1], "\""), nil
		}
	}

	return "", fmt.Errorf("unable to determine Linux distribution")
}

func listIptables() ([]string, error) {
	cmd := exec.Command("iptables", "-L")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	return lines, nil
}

func listFirewalld() ([]string, error) {
	cmd := exec.Command("firewall-cmd", "--list-all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	return lines, nil
}

func listUfw() ([]string, error) {
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
	cmd := exec.Command("ufw", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func addIpTablesRule(rule string) error {
	cmd := exec.Command("iptables", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func addFirewalldRule(rule string) error {
	cmd := exec.Command("firewall-cmd", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeUfwRule(rule string) error {
	cmd := exec.Command("ufw", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeIpTablesRule(rule string) error {
	cmd := exec.Command("iptables", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeFirewalldRule(rule string) error {
	cmd := exec.Command("firewall-cmd", rule)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func getOSName() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			name := strings.TrimSpace(line[len("PRETTY_NAME="):])
			if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
				name = name[1 : len(name)-1]
			}
			return name, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("OS name not found")
}

func getKernelVersion() (string, error) {
	kernelVersion := exec.Command("uname", "-r")
	out, err := kernelVersion.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// getUptime returns the system uptime as a time.Duration.
func getUptime() (string, error) {
	cmd := exec.Command("uptime", "-p")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
