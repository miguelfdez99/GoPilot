package backend

import (
	"fmt"
	"os/exec"
)

func (b *Backend) ListPackages() []string {

	distribution, err := getLinuxDistribution()
	if err != nil {
		b.logger.WithError(err).Error("Failed to get Linux distribution")
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
		b.logger.WithError(err).Error("Failed to list packages")
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
		b.logger.WithError(err).Error("Failed to get Linux distribution")
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
		b.logger.WithError(err).Error("Failed to remove package")
		return err
	}

	b.logger.Infof("Package %s removed successfully", pkgName)

	return nil
}

func (b *Backend) InstallPackage(pkgName string) error {
	distribution, err := getLinuxDistribution()
	if err != nil {
		b.logger.WithError(err).Error("Failed to get Linux distribution")
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
		b.logger.WithError(err).Error("Failed to install package")
		return err
	}

	b.logger.Infof("Package %s installed successfully", pkgName)

	return nil
}

func (b *Backend) GetDistribution() string {
	distribution, err := getLinuxDistribution()
	if err != nil {
		b.logger.WithError(err).Error("Failed to get distribution")
		return ""
	}
	return distribution
}
