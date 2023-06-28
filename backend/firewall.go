package backend

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var distributionsMap = map[string]func(string) error{
	"fedora": addFirewalldRule,
	"rhel":   addFirewalldRule,
	"centos": addFirewalldRule,
	"ubuntu": addUfwRule,
	"debian": addUfwRule,
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

func (b *Backend) ListFirewallRules() ([]string, error) {
	distribution, err := getLinuxDistribution()
	if err != nil {
		return nil, err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return listFirewalld()
	} else if distribution == "ubuntu" || distribution == "debian" {
		return listUfw()
	}

	return listIptables()
}

func (b *Backend) AddFirewallRule(rule string) error {
	if !isValidRule(rule) {
		return errors.New("invalid rule")
	}

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	addFunc, exists := distributionsMap[distribution]
	if !exists {
		addFunc = addIpTablesRule
	}

	return addFunc(rule)
}

func (b *Backend) RemoveFirewallRule(rule string) error {
	if !isValidRule(rule) {
		return errors.New("invalid rule")
	}

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	removeFunc, exists := distributionsMap[distribution]
	if !exists {
		removeFunc = removeIpTablesRule
	}

	return removeFunc(rule)
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
	args := append([]string{"ufw"}, strings.Fields(rule)...)
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func addIpTablesRule(rule string) error {
	args := append([]string{"iptables"}, strings.Fields(rule)...)
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func addFirewalldRule(rule string) error {
	args := append([]string{"firewall-cmd"}, strings.Fields(rule)...)
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

func removeIpTablesRule(rule string) error {
	args := append([]string{"iptables", "-D"}, strings.Fields(rule)...)
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeFirewalldRule(rule string) error {
	args := append([]string{"firewall-cmd", "--remove"}, strings.Fields(rule)...)
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
