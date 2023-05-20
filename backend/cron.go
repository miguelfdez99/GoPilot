package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type CronJob struct {
	Schedule string
	Command  string
}

func (b *Backend) ListCronJobs() ([]CronJob, error) {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		b.logger.Error("Failed to list cron jobs")
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	jobs := []CronJob{}

	specialSchedulePattern := regexp.MustCompile(`^\s*@[^ \t]+\s+.*$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		if specialSchedulePattern.MatchString(line) {
			parts := strings.Fields(line)

			if len(parts) < 2 {
				continue
			}

			job := CronJob{
				Schedule: parts[0],
				Command:  strings.Join(parts[1:], " "),
			}
			jobs = append(jobs, job)
		} else {
			parts := strings.SplitN(line, " ", 6)
			if len(parts) < 6 {
				continue
			}

			job := CronJob{
				Schedule: strings.Join(parts[:5], " "),
				Command:  strings.Join(parts[5:], " "),
			}
			jobs = append(jobs, job)
		}

	}

	return jobs, nil
}

func (b *Backend) RemoveAllCronJobs() error {
	cmd := exec.Command("crontab", "-r")
	err := cmd.Run()
	if err != nil {
		b.logger.Error("Failed to remove all cron jobs")
		return err
	}
	b.logger.Info("All cron jobs removed successfully")
	return nil
}

func (b *Backend) RemoveCronJob(job string) error {

	jobs, err := b.ListCronJobs()
	if err != nil {
		b.logger.Error("Failed to list cron jobs for removal")
		return err
	}

	index := -1
	for i, j := range jobs {
		if j.Command == job {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("cron job not found: %s", job)
	}

	jobs = append(jobs[:index], jobs[index+1:]...)

	err = b.updateCronTab(jobs)
	if err != nil {
		b.logger.Error("Failed to update cron tab")
		return err
	}

	b.logger.Info(fmt.Sprintf("Cron job removed successfully: %s", job))
	return nil
}

func (b *Backend) updateCronTab(jobs []CronJob) error {

	var cronTabBuffer strings.Builder
	for _, job := range jobs {
		cronTabBuffer.WriteString(fmt.Sprintf("%s %s\n", job.Schedule, job.Command))
	}

	cmd := exec.Command("crontab", "-")
	cmd.Stdin = strings.NewReader(cronTabBuffer.String())
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (b *Backend) AddCronJob(schedule string, command string) error {
	tmpfile, err := ioutil.TempFile("", "crontab")
	if err != nil {
		b.logger.Error("Failed to create temp file")
		return err
	}
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	_, err = tmpfile.Write(output)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(tmpfile, "%s %s\n", schedule, command)
	if err != nil {
		return err
	}

	cmd = exec.Command("crontab", tmpfile.Name())
	err = cmd.Run()
	if err != nil {
		b.logger.Error("Failed to run cron tab add command")
		return err
	}

	b.logger.Info(fmt.Sprintf("Cron job added successfully: %s %s", schedule, command))
	return nil
}
