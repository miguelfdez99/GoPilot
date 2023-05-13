package backend

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func (b *Backend) GetLSCPU() string {
	cmd := exec.Command("lscpu")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	lines := strings.Split(string(out), "\n")
	var architecture, cpus, modelName, threadPerCore, corePerSocket, socket, cpuModes string

	for _, line := range lines {
		if strings.HasPrefix(line, "Architecture:") {
			architecture = strings.TrimSpace(line[len("Architecture:"):])
		} else if strings.HasPrefix(line, "CPU(s):") {
			cpus = strings.TrimSpace(line[len("CPU(s):"):])
		} else if strings.HasPrefix(line, "Model name:") {
			modelName = strings.TrimSpace(line[len("Model name:"):])
		} else if strings.HasPrefix(line, "Thread(s) per core:") {
			threadPerCore = strings.TrimSpace(line[len("Thread(s) per core:"):])
		} else if strings.HasPrefix(line, "Core(s) per socket:") {
			corePerSocket = strings.TrimSpace(line[len("Core(s) per socket:"):])
		} else if strings.HasPrefix(line, "Socket(s):") {
			socket = strings.TrimSpace(line[len("Socket(s):"):])
		} else if strings.HasPrefix(line, "CPU op-mode(s):") {
			cpuModes = strings.TrimSpace(line[len("CPU op-mode(s):"):])
		}
	}

	data := map[string]string{
		"architecture":  architecture,
		"cpus":          cpus,
		"modelName":     modelName,
		"threadPerCore": threadPerCore,
		"corePerSocket": corePerSocket,
		"socket":        socket,
		"cpuModes":      cpuModes,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(jsonData))
	return string(jsonData)
}

func (b *Backend) GetSystemInfo() (string, error) {
	osName, err := getOSName()
	if err != nil {
		return "", fmt.Errorf("error obtaining operating system name: %v", err)
	}

	kernelVer, err := getKernelVersion()
	if err != nil {
		return "", fmt.Errorf("error obtaining kernel version: %v", err)
	}

	uptime, err := getUptime()
	if err != nil {
		return "", fmt.Errorf("error obtaining system uptime: %v", err)
	}

	data := map[string]string{
		"osName":    osName,
		"kernelVer": kernelVer,
		"uptime":    uptime,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error converting system info to JSON: %v", err)
	}

	return string(jsonData), nil
}

func (b *Backend) GetCPUUsage() ([]float64, error) {
	// Get percpu usage
	percpuUsage, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}

	return percpuUsage, nil
}
