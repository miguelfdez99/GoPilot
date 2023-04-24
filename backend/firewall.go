package backend

func (b *Backend) ListFirewallRules() ([]string, error) {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return nil, err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return listFirewalld()
	} else if distribution == "ubuntu" {
		return listUfw()
	}

	return listIptables()
}

func (b *Backend) AddFirewallRule(rule string) error {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return addFirewalldRule(rule)
	} else if distribution == "ubuntu" {
		return addUfwRule(rule)
	}

	return addIpTablesRule(rule)
}

func (b *Backend) RemoveFirewallRule(rule string) error {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return removeFirewalldRule(rule)
	} else if distribution == "ubuntu" {
		return removeUfwRule(rule)
	}

	return removeIpTablesRule(rule)
}
