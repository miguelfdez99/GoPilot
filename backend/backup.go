package backend

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type BackupOptions struct {
	SourceDir    string
	DestDir      string
	Exclude      []string
	CompressData bool
	LinksOption  string
	Verify       bool
	CompressFile bool
}

func (b *Backend) Backup(options BackupOptions) (string, error) {
	args := []string{"-av", "--delete"}

	if options.CompressData {
		args = append(args, "-z")
	}

	if options.LinksOption != "" {
		args = append(args, options.LinksOption)
	}

	for _, path := range options.Exclude {
		args = append(args, fmt.Sprintf("--exclude=%s", path))
	}

	args = append(args, options.SourceDir, options.DestDir)

	cmdString := fmt.Sprintf("rsync %s", strings.Join(args, " "))

	cmd := exec.Command("rsync", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error creating backup:", err)
		return "", fmt.Errorf("error creating backup: %s", string(output))
	}

	b.logger.Info(fmt.Sprintf("Backup created successfully:\n%s", string(output)))

	if options.Verify {

		sourceBase := filepath.Base(options.SourceDir)
		dirDest := filepath.Join(options.DestDir, sourceBase)

		cmd = exec.Command("diff", "-rq", options.SourceDir, dirDest)
		output, err = cmd.CombinedOutput()

		if err != nil {
			log.Println("Backup verification failed:", err)
			return "", fmt.Errorf("backup verification failed: %s", string(output))
		}

		b.logger.Info(fmt.Sprintf("Backup verified successfully:\n%s", string(output)))
	}

	if options.CompressFile {
		cmd = exec.Command("tar", "-czf", options.DestDir+".tar.gz", "-C", options.DestDir, ".")
		output, err = cmd.CombinedOutput()

		if err != nil {
			log.Println("Error compressing backup:", err)
			return "", fmt.Errorf("error compressing backup: %s", string(output))
		}

		b.logger.Info(fmt.Sprintf("Backup compressed successfully:\n%s", string(output)))
	}

	return cmdString, nil
}

func (b *Backend) ScheduleBackup(options BackupOptions, schedule string) error {
	cmdString, err := b.Backup(options)
	if err != nil {
		return err
	}

	b.logger.Info(fmt.Sprint("Scheduling backup: ", cmdString))

	return b.AddCronJob(schedule, cmdString)
}
