package backend

import (
	"errors"
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
