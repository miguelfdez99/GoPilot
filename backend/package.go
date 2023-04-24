package backend

import (
	"fmt"
	"os/exec"
)

func (b *Backend) ListPackages() []string {

	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "list", "--installed")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "list", "installed")
	case "arch":
		cmd = exec.Command("pacman", "-Q")
	default:
		fmt.Println("Unsupported Linux distribution")
		return nil
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	packageNames := ExtractFirstParams(string(out))
	if distribution != "arch" {
		packageNames = packageNames[1:]
	}

	return packageNames
}

func (b *Backend) RemovePackage(pkgName string) error {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "remove", pkgName, "-y")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "remove", pkgName, "-y")
	case "arch":
		cmd = exec.Command("pacman", "-R", pkgName, "--noconfirm")
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("Package %s removed successfully", pkgName)

	return nil
}

func (b *Backend) InstallPackage(pkgName string) error {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "install", pkgName, "-y")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "install", pkgName, "-y")
	case "arch":
		cmd = exec.Command("pacman", "-S", pkgName, "--noconfirm")
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("Package %s installed successfully", pkgName)

	return nil
}

func (b *Backend) GetDistribution() string {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(distribution)
	return distribution
}
