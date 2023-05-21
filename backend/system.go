package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
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
	return string(jsonData)
}

type SystemInfo struct {
	OsName          string `json:"osName"`
	KernelVer       string `json:"kernelVer"`
	Uptime          string `json:"uptime"`
	Hostname        string `json:"hostname"`
	DesktopEnv      string `json:"desktopEnv"`
	CurrentUsername string `json:"currentUsername"`
	Memory          string `json:"memory"`
}

func (b *Backend) GetSystemInfo() (string, error) {
	var sysInfo SystemInfo
	var err error

	if sysInfo.OsName, err = getOSName(); err != nil {
		return "", fmt.Errorf("error obtaining operating system name: %v", err)
	}

	if sysInfo.KernelVer, err = getKernelVersion(); err != nil {
		return "", fmt.Errorf("error obtaining kernel version: %v", err)
	}

	if sysInfo.Uptime, err = getUptime(); err != nil {
		return "", fmt.Errorf("error obtaining system uptime: %v", err)
	}

	if sysInfo.Hostname, err = os.Hostname(); err != nil {
		return "", fmt.Errorf("error obtaining system hostname: %v", err)
	}

	sysInfo.DesktopEnv = getDesktopEnv()

	if currentUser, err := user.Current(); err == nil {
		sysInfo.CurrentUsername = currentUser.Username
	} else {
		return "", fmt.Errorf("error obtaining current user: %v", err)
	}

	if sysInfo.Memory, err = getMemory(); err != nil {
		return "", fmt.Errorf("error obtaining memory info: %v", err)
	}

	jsonData, err := json.Marshal(sysInfo)
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

type MemoryUsage struct {
	RAM  float64 `json:"ram"`
	Swap float64 `json:"swap"`
}

func (b *Backend) GetMemoryUsage() (*MemoryUsage, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	s, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	ramUsage := 100.0 * (float64(v.Total) - float64(v.Available)) / float64(v.Total)
	swapUsage := 100.0 * (float64(s.Total) - float64(s.Free)) / float64(s.Total)

	return &MemoryUsage{
		RAM:  ramUsage,
		Swap: swapUsage,
	}, nil
}
