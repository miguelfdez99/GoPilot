package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func (b *Backend) DeleteTempFiles(dirPath string, expireDays int) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			b.logger.Error(fmt.Sprintf("Failure accessing a path %q: %v", path, err))
			return err
		}

		if time.Since(info.ModTime()).Hours()/24 > float64(expireDays) {
			err := os.RemoveAll(path)
			if err != nil {
				b.logger.Error(fmt.Sprintf("Error deleting file %q: %v", path, err))
			} else {
				b.logger.Info(fmt.Sprintf("Deleted file %q", path))
			}
		}
		return nil
	})

	if err != nil {
		b.logger.Error(fmt.Sprintf("Error walking the path %q: %v", dirPath, err))
		return
	}

	b.logger.Info(fmt.Sprintf("Cleanup of directory %q completed", dirPath))
}

func (b *Backend) RemoveUnusedPackages() {

	distribution, err := getLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get Linux distribution")
	}

	switch distribution {
	case "ubuntu", "debian":
		exec.Command("apt", "autoremove", "-y").Run()
	case "fedora", "centos", "rhel":
		exec.Command("dnf", "autoremove", "-y").Run()
	case "arch":
		exec.Command("pacman", "-Rns", "$(pacman -Qtdq)", "--noconfirm").Run()
	default:
		fmt.Println("Unsupported Linux distribution")
	}
	b.logger.Info("Unused packages removed")
}

func (b *Backend) CleanCachePackages() {
	distribution, err := getLinuxDistribution()
	if err != nil {
		b.logger.Error("Failed to get Linux distribution")
	}

	switch distribution {
	case "ubuntu", "debian":
		exec.Command("apt", "clean").Run()
	case "fedora", "centos", "rhel":
		exec.Command("dnf", "clean", "all").Run()
	case "arch":
		exec.Command("pacman", "-Sc", "--noconfirm").Run()
	default:
		fmt.Println("Unsupported Linux distribution")
	}
	b.logger.Info("Cache packages cleaned")
}

func (b *Backend) DuplicatedFiles(dirPath string) {
	output, err := exec.Command("fdupes", "-r", dirPath).Output()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to execute fdupes on %q: %v", dirPath, err))
		return
	}
	files := strings.Split(string(output), "\n")
	b.logger.Info(fmt.Sprintf("Duplicated files in %q: %v", dirPath, files))
	for i := 1; i < len(files); i++ {
		if files[i] == "" {
			i++
			continue
		}
		err := os.Remove(files[i])
		if err != nil {
			b.logger.Error(fmt.Sprintf("Error deleting file %q: %v", files[i], err))
		} else {
			b.logger.Info(fmt.Sprintf("Deleted file %q", files[i]))
		}
	}
}

func (b *Backend) RemoveOldLogs(days string) {
	exec.Command("journalctl", "--vacuum-time="+days).Run()
	b.logger.Info(fmt.Sprintf("Removed logs older than %q days", days))
}
