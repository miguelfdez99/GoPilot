package backend

import (
	"fmt"
	"goPilot/backend/utils"
	"os/exec"
)

func (b *Backend) ListPackages() []string {

	distribution, err := utils.GetLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get Linux distribution")
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
	case "opensuse":
		cmd = exec.Command("zypper", "se", "--installed-only")
	default:
		fmt.Println("Unsupported Linux distribution")
		return nil
	}

	out, err := cmd.Output()
	if err != nil {
		b.logger.Error("Failed to list packages")
		return nil
	}
	packageNames := utils.ExtractFirstParams(string(out))
	if distribution != "arch" {
		packageNames = packageNames[1:]
	}

	return packageNames
}

func (b *Backend) RemovePackage(pkgName string) error {
	distribution, err := utils.GetLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get Linux distribution")
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
	case "opensuse":
		cmd = exec.Command("zypper", "--non-interactive", "rm", "--clean-deps", pkgName)
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		b.logger.Error("Failed to remove package")
		return err
	}

	b.logger.Info(fmt.Sprintf("Package %s removed successfully", pkgName))

	return nil
}

func (b *Backend) InstallPackage(pkgName string) error {
	distribution, err := utils.GetLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get Linux distribution")
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
	case "opensuse":
		cmd = exec.Command("zypper", "--non-interactive", "in", pkgName)
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		b.logger.Error("Failed to install package")
		return err
	}

	b.logger.Info(fmt.Sprintf("Package %s installed successfully", pkgName))

	return nil
}

func (b *Backend) GetDistribution() string {
	distribution, err := utils.GetLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get distribution")
		return ""
	}
	return distribution
}
