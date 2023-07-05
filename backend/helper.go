package backend

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func (b *Backend) CheckAdmin() bool {
	uid := os.Getuid()
	if uid != 0 {
		fmt.Println("You are not an admin")
		return false
	} else {
		fmt.Println("You are an admin")
		return true
	}
}

func ExtractFirstParams(input string) []string {
	distribution, err := getLinuxDistribution()
	if err != nil {
		return nil
	}

	lines := strings.Split(input, "\n")
	params := make([]string, len(lines))
	for i, line := range lines {
		if distribution == "debian" || distribution == "ubuntu" {
			fields := strings.Split(line, "/")
			params[i] = fields[0]
		} else if distribution == "fedora" || distribution == "centos" || distribution == "rhel" {
			fields := strings.Split(line, ".")
			params[i] = fields[0]
		} else if distribution == "arch" {
			fields := strings.Split(line, " ")
			params[i] = fields[0]
		}
	}
	return params
}

func getLinuxDistribution() (string, error) {
	fileContent, err := ioutil.ReadFile("/etc/os-release")
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
			return strings.Trim(parts[1], "\""), nil
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

func getOSName() (string, error) {
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

func getKernelVersion() (string, error) {
	kernelVersion := exec.Command("uname", "-r")
	out, err := kernelVersion.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func getUptime() (string, error) {
	cmd := exec.Command("uptime", "-p")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func getMemory() (string, error) {
	cmd := exec.Command("free", "-h")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return "", fmt.Errorf("unable to parse memory info")
	}

	return strings.TrimSpace(lines[1]), nil
}

// func getDesktopEnv() string {
// 	if desktopEnv := os.Getenv("XDG_CURRENT_DESKTOP"); desktopEnv != "" {
// 		return strings.Split(desktopEnv, ":")[1]
// 	}

// 	if desktopEnv := os.Getenv("DESKTOP_SESSION"); desktopEnv != "" {
// 		return desktopEnv
// 	}

// 	if desktopEnv := os.Getenv("GDMSESSION"); desktopEnv != "" {
// 		return desktopEnv
// 	}

// 	if desktopEnv, err := getDesktopEnvFromXprop(); err == nil {
// 		return desktopEnv
// 	}

// 	return "Not Found"
// }

// func getDesktopEnvFromXprop() (string, error) {
// 	cmd := exec.Command("xprop", "-root")
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, line := range strings.Split(string(output), "\n") {
// 		if strings.Contains(line, "_DT_SAVE_MODE") {
// 			return "Xfce", nil
// 		} else if strings.Contains(line, "MUFFIN") {
// 			return "Cinnamon", nil
// 		}
// 	}

// 	return "", fmt.Errorf("could not determine desktop environment from xprop")
// }

func (b *Backend) CommandExists(cmd string) (bool, string) {
	_, err := exec.LookPath(cmd)
	if err != nil {
		b.logger.Error(fmt.Sprintf("%s command is not available. Error: %v", cmd, err))
		return false, fmt.Sprintf("%s command is not available", cmd)
	}
	return true, ""
}
