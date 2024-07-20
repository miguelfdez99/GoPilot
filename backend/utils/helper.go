package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ExtractFirstParams(input string) []string {
	distribution, err := GetLinuxDistribution()
	if err != nil {
		return nil
	}

	lines := strings.Split(input, "\n")
	var params []string
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch distribution {
		case "debian", "ubuntu":
			parts := strings.Split(fields[0], "/")
			params = append(params, parts[0])
		case "fedora", "centos", "rhel":
			parts := strings.Split(fields[0], ".")
			params = append(params, parts[0])
		case "arch":
			params = append(params, fields[0])
		case "opensuse":
			if fields[0] == "i" && fields[1] == "|" {
				params = append(params, fields[2])
			}
		}
	}
	return params
}

func GetLinuxDistribution() (string, error) {
	fileContent, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(fileContent))
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID_LIKE" {
			value := strings.Split(strings.Trim(parts[1], "\""), " ")
			return value[0], nil
		}
	}

	// If ID_LIKE is not present, try ID(Fedora does not have ID_LIKE)
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID" {
			return strings.Trim(parts[1], "\""), nil
		}
	}

	return "", fmt.Errorf("unable to determine Linux distribution")
}

func GetOSName() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			name := strings.TrimSpace(line[len("PRETTY_NAME="):])
			if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
				name = name[1 : len(name)-1]
			}
			return name, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("OS name not found")
}

func GetKernelVersion() (string, error) {
	kernelVersion := exec.Command("uname", "-r")
	out, err := kernelVersion.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetUptime() (string, error) {
	cmd := exec.Command("uptime")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	data := strings.TrimSpace(string(out))

	re := regexp.MustCompile(`up\s+((\d+\s+days?,\s+)?\d+:\d+|\d+\s+min)`)
	matches := re.FindStringSubmatch(data)

	if len(matches) < 2 {
		return "", nil
	}

	return matches[1], nil
}

func GetMemory() (string, error) {
	cmd := exec.Command("free", "-h")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return "", fmt.Errorf("unable to parse memory info")
	}

	parts := strings.Fields(lines[1])
	if len(parts) < 3 {
		return "", fmt.Errorf("unexpected format of memory info")
	}

	totalMemory := parts[1]
	usedMemory := parts[2]

	memoryInfo := fmt.Sprintf("Total: %s, Used: %s", totalMemory, usedMemory)
	return memoryInfo, nil
}

func GetDesktopEnv() (string, error) {
	if desktopEnv := os.Getenv("XDG_CURRENT_DESKTOP"); desktopEnv != "" {
		splitEnv := strings.Split(desktopEnv, ":")
		if len(splitEnv) > 1 {
			return splitEnv[1], nil
		} else {
			return splitEnv[0], nil
		}
	}

	if desktopEnv := os.Getenv("DESKTOP_SESSION"); desktopEnv != "" {
		return desktopEnv, nil
	}

	if desktopEnv := os.Getenv("GDMSESSION"); desktopEnv != "" {
		return desktopEnv, nil
	}

	if desktopEnv, err := GetDesktopEnvFromXprop(); err == nil {
		return desktopEnv, nil
	}

	return "Not Found", nil
}

func GetDesktopEnvFromXprop() (string, error) {
	cmd := exec.Command("xprop", "-root")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "_DT_SAVE_MODE") {
			return "Xfce", nil
		} else if strings.Contains(line, "MUFFIN") {
			return "Cinnamon", nil
		}
	}

	return "", fmt.Errorf("could not determine desktop environment from xprop")
}
