package backend

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Service struct {
	Name          string
	StartupStatus string
}

type RunningService struct {
	Name        string
	ActiveState string
	Description string
}

func (b *Backend) GetAllServices() ([]Service, error) {
	cmd := exec.Command("systemctl", "list-unit-files", "--type=service", "--state=enabled", "--state=disabled", "--no-pager", "--no-legend", "--plain", "--all")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprint("Failed to list services: ", err))
		return nil, err
	}

	var services []Service
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 3 && !strings.Contains(fields[0], "@") {
			services = append(services, Service{
				Name:          fields[0],
				StartupStatus: fields[1],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		b.logger.Error(fmt.Sprint("Failed to scan command output: ", err))
		return nil, err
	}

	return services, nil
}

func (b *Backend) GetRunningServices() ([]RunningService, error) {
	cmd := exec.Command("systemctl", "--type=service", "--all", "--no-legend")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to list running services: %v", err))
		return nil, err
	}

	var services []RunningService
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "@") {
			continue
		}

		if strings.Contains(line, "●") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 4 {
			description := strings.Join(fields[4:], " ")

			services = append(services, RunningService{
				Name:        fields[0],
				ActiveState: fields[2],
				Description: description,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		b.logger.Error(fmt.Sprintf("Failed to scan command output: %v", err))
		return nil, err
	}

	fmt.Println(services)

	return services, nil
}

func (b *Backend) EnableService(service string) {
	cmd := exec.Command("systemctl", "enable", service)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprint("Failed to enable service: ", service, "Error: ", err))
	}
	b.logger.Info(fmt.Sprint("Enabled service: ", service))
}

func (b *Backend) DisableService(service string) {
	cmd := exec.Command("systemctl", "disable", service)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprint("Failed to disable service: ", service, "Error: ", err))
	}
	b.logger.Info(fmt.Sprint("Disabled service: ", service))
}

func (b *Backend) StartService(service string) {
	cmd := exec.Command("systemctl", "start", service)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprint("Failed to start service: ", service, "Error: ", err))
	}
	b.logger.Info(fmt.Sprint("Started service: ", service))
}

func (b *Backend) StopService(service string) {
	cmd := exec.Command("systemctl", "stop", service)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprint("Failed to stop service: ", service, "Error: ", err))
	}
	b.logger.Info(fmt.Sprint("Stopped service: ", service))
}
