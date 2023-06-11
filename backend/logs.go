package backend

import (
	"bufio"
	"fmt"
	"os/exec"
)

func (b *Backend) GetLogs(logType string) ([]string, error) {
	var command *exec.Cmd
	switch logType {
	case "all":
		command = exec.Command("journalctl", "--no-pager")
	case "system":
		command = exec.Command("journalctl", "--no-pager", "-k")
	case "application":
		command = exec.Command("journalctl", "--no-pager")
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
