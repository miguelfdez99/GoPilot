package backend

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"
)

type BackupOptions struct {
	SourceDir    string
	DestDir      string
	Exclude      []string
	CompressData bool
	LinksOption  string
	Verify       bool
	Schedule     string
	CompressFile bool
}

func (b *Backend) Backup(options BackupOptions) error {
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

	cmd := exec.Command("rsync", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error creating backup:", err)
		return fmt.Errorf("error creating backup: %s", string(output))
	}

	log.Println("Backup created successfully:\n", string(output))

	if options.Verify {

		sourceBase := filepath.Base(options.SourceDir)
		dirDest := filepath.Join(options.DestDir, sourceBase)

		cmd = exec.Command("diff", "-rq", options.SourceDir, dirDest)
		output, err = cmd.CombinedOutput()

		if err != nil {
			log.Println("Backup verification failed:", err)
			return fmt.Errorf("backup verification failed: %s", string(output))
		}

		log.Println("Backup verified successfully:\n", string(output))
	}

	if options.Schedule != "" {
		duration, err := time.ParseDuration(options.Schedule)
		if err != nil {
			log.Println("Invalid schedule duration:", err)
			return fmt.Errorf("invalid schedule duration: %s", options.Schedule)
		}

		time.AfterFunc(duration, func() {
			err := b.Backup(options)
			if err != nil {
				log.Println("Scheduled backup failed:", err)
			}
		})

		log.Println("Next backup scheduled in", duration)
	}

	if options.CompressFile {
		cmd = exec.Command("tar", "-czf", options.DestDir+".tar.gz", "-C", options.DestDir, ".")
		output, err = cmd.CombinedOutput()

		if err != nil {
			log.Println("Error compressing backup:", err)
			return fmt.Errorf("error compressing backup: %s", string(output))
		}

		log.Println("Backup compressed successfully:\n", string(output))
	}

	return nil
}
