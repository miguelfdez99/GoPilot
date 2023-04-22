package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

type User struct {
	Username      string
	Password      string
	UID           int
	GID           int
	HomeDirectory string
	Shell         string
	Groups        []string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

///////////////////////////-----------------///////////////////////////

func (a *App) CreateUser(user User) error {
	cmd := exec.Command("useradd",
		"-m",
		"-s", user.Shell,
		"-u", fmt.Sprint(user.UID),
		"-g", fmt.Sprint(user.GID),
		"-p", user.Password,
		user.Username,
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (a *App) DeleteUser(username string, removeHomeDir bool, forceDelete bool) error {
	args := []string{}
	if removeHomeDir {
		args = append(args, "-r")
	}
	if forceDelete {
		args = append(args, "-f")
	}
	args = append(args, username)
	cmd := exec.Command("userdel", args...)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error deleting user %s: %s", username, err)
	}
	return nil
}

func (a *App) ModifyUser(username string, userPtr *User) error {
	cmd := exec.Command("usermod", username)

	if userPtr != nil {
		if userPtr.Shell != "" {
			cmd.Args = append(cmd.Args, "-s", userPtr.Shell)
		}
		if userPtr.UID != 0 {
			cmd.Args = append(cmd.Args, "-u", fmt.Sprint(userPtr.UID))
		}
		if userPtr.GID != 0 {
			cmd.Args = append(cmd.Args, "-g", fmt.Sprint(userPtr.GID))
		}
		if userPtr.Password != "" {
			cmd.Args = append(cmd.Args, "-p", userPtr.Password)
		}
		if userPtr.HomeDirectory != "" {
			cmd.Args = append(cmd.Args, "-d", userPtr.HomeDirectory, "-m")
		}
		if len(userPtr.Groups) > 0 {
			cmd.Args = append(cmd.Args, "-aG", strings.Join(userPtr.Groups, ","))
		}
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to modify user: %v", err)
	}
	return nil
}

func (a *App) CreateGroup(name string, gid *int) error {
	cmdArgs := []string{"groupadd", name}
	if gid != nil {
		cmdArgs = append(cmdArgs, "-g", strconv.Itoa(*gid))
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ModifyGroup(name string, gid int) error {
	cmd := exec.Command("groupmod", "-g", strconv.Itoa(gid), name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) DeleteGroup(name string) error {
	cmd := exec.Command("groupdel", name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

///////////////////////////-----------------///////////////////////////

func (a *App) CheckAdmin() bool {
	uid := os.Getuid()
	if uid != 0 {
		fmt.Println("You are not an admin")
		return false
	} else {
		fmt.Println("You are an admin")
		return true
	}
}

///////////////////////////-----------------///////////////////////////

func (a *App) ListPackages() []string {

	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "list", "--installed")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "list", "installed")
	case "arch":
		cmd = exec.Command("pacman", "-Q")
	default:
		fmt.Println("Unsupported Linux distribution")
		return nil
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	packageNames := ExtractFirstParams(string(out))
	if distribution != "arch" {
		packageNames = packageNames[1:]
	}

	return packageNames
}

func (a *App) RemovePackage(pkgName string) error {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "remove", pkgName, "-y")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "remove", pkgName, "-y")
	case "arch":
		cmd = exec.Command("pacman", "-R", pkgName, "--noconfirm")
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("Package %s removed successfully", pkgName)

	return nil
}

func (a *App) InstallPackage(pkgName string) error {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var cmd *exec.Cmd

	switch distribution {
	case "ubuntu", "debian":
		cmd = exec.Command("apt", "install", pkgName, "-y")
	case "fedora", "centos", "rhel":
		cmd = exec.Command("dnf", "install", pkgName, "-y")
	case "arch":
		cmd = exec.Command("pacman", "-S", pkgName, "--noconfirm")
	default:
		return fmt.Errorf("unsupported Linux distribution: %s", distribution)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("Package %s installed successfully", pkgName)

	return nil
}

func (a *App) GetDistribution() string {
	distribution, err := getLinuxDistribution()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(distribution)
	return distribution
}

///////////////////////////-----------------///////////////////////////

func (a *App) GetLSCPU() string {
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

func (a *App) GetSystemInfo() (string, error) {
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

///////////////////////////-----------------///////////////////////////

type CronJob struct {
	Schedule string
	Command  string
}

func (a *App) ListCronJobs() ([]CronJob, error) {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	jobs := []CronJob{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
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

	return jobs, nil
}

func (a *App) RemoveAllCronJobs() error {
	cmd := exec.Command("crontab", "-r")
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("All cron jobs removed successfully")
	return nil
}

func (a *App) AddCronJob(schedule string, command string) error {
	tmpfile, err := ioutil.TempFile("", "crontab")
	if err != nil {
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
		return err
	}

	return nil
}

///////////////////////////-----------------///////////////////////////

func (a *App) ListFirewallRules() ([]string, error) {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return nil, err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return listFirewalld()
	} else if distribution == "ubuntu" {
		return listUfw()
	}

	return listIptables()
}

func (a *App) AddFirewallRule(rule string) error {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return addFirewalldRule(rule)
	} else if distribution == "ubuntu" {
		return addUfwRule(rule)
	}

	return addIpTablesRule(rule)
}

func (a *App) RemoveFirewallRule(rule string) error {

	distribution, err := getLinuxDistribution()
	if err != nil {
		return err
	}

	if distribution == "fedora" || distribution == "rhel" || distribution == "centos" {
		return removeFirewalldRule(rule)
	} else if distribution == "ubuntu" {
		return removeUfwRule(rule)
	}

	return removeIpTablesRule(rule)
}
