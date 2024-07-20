package backend

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func (b *Backend) Backup(options BackupOptions) (string, error) {
	backupFileName := filepath.Base(options.SourceDir) + ".tar.gz"
	compressedFilePath := filepath.Join(options.DestDir, backupFileName)
	args := []string{"-czf", compressedFilePath}
	args = append(args, "-C", filepath.Dir(options.SourceDir), filepath.Base(options.SourceDir))

	cmd := exec.Command("tar", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error compressing backup:", err)
		return "", fmt.Errorf("error compressing backup: %s", string(output))
	}

	b.logger.Info(fmt.Sprintf("Backup compressed successfully:\n%s", string(output)))

	return compressedFilePath, nil
}
