package backend

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func (b *Backend) GetLogs(logType string, boot int) ([]string, error) {
	bootFlag := fmt.Sprintf("-b %d", boot)
	var command *exec.Cmd
	switch logType {
	case "all":
		command = exec.Command("journalctl", "--no-pager", bootFlag)
	case "system":
		command = exec.Command("journalctl", "--no-pager", "-k", bootFlag)
	case "hardware":
		command = exec.Command("sh", "-c", "journalctl --no-pager -k "+bootFlag+" | grep -iE 'usb|pci|hdmi|hid|eth|wlan|bluetooth'")
	case "application":
		command = exec.Command("sh", "-c", fmt.Sprintf("journalctl --no-pager %s | grep -v 'kernel\\|systemd'", bootFlag))
	case "important":
		command = exec.Command("journalctl", "--no-pager", "-p", "err", bootFlag)
	default:
		return nil, fmt.Errorf("unknown log type: %s", logType)
	}

	outputPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := command.Start(); err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(outputPipe)
	logs := make(chan string)

	go func() {
		for scanner.Scan() {
			logs <- scanner.Text()
		}
		close(logs)
	}()

	var logLines []string
	for logLine := range logs {
		logLines = append(logLines, logLine)
	}

	if err := command.Wait(); err != nil {
		return nil, err
	}

	//fmt.Println(logLines)
	return logLines, nil
}

func (b *Backend) ExportLogs(logType string, filepath string) error {
	var command *exec.Cmd

	switch logType {
	case "all":
		command = exec.Command("journalctl", "--no-pager")
	case "system":
		command = exec.Command("journalctl", "--no-pager", "-k")
	case "hardware":
		command = exec.Command("journalctl", "--no-pager", "_SYSLOG_IDENTIFIER=kernel")
	case "application":
		command = exec.Command("journalctl", "--no-pager")
	case "important":
		command = exec.Command("journalctl", "--no-pager", "-p", "err")
	default:
		return nil
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	command.Stdout = file

	if err := command.Start(); err != nil {
		return err
	}

	if err := command.Wait(); err != nil {
		return err
	}

	return nil
}
