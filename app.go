package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
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

func (a *App) PrintUsers() {
	fmt.Println("User ID:", os.Getuid())
	fmt.Println("Group ID:", os.Getgid())
	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Group IDs:", groups)
}

func (a *App) CreateUser() error {

	distribution, err := getLinuxDistribution()

	var cmd *exec.Cmd
	if distribution == "ubuntu" || distribution == "debian" {
		cmd = exec.Command("adduser", "testuser")
	} else {
		cmd = exec.Command("useradd", "testuser")
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	if stderr.Len() > 0 && stderr.String() == "adduser: The user `testuser' already exists.\n" {
		fmt.Printf("error creating user: %v\nStderr output: %s", err, stderr.String())
	}

	if err != nil {
		return err
	}
	fmt.Printf("User %s created successfully\n", "testuser")
	return nil
}

func (a *App) CreateUser2(user User) error {
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

func (a *App) CreateUserWithDir(username string) error {
	cmd := exec.Command("adduser", "--home", "/home/"+username, "--disabled-password", "--gecos", "", username)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if stderr.Len() > 0 && strings.Contains(stderr.String(), "The user `"+username+"' already exists.") {
		fmt.Printf("error creating user: %v\nStderr output: %s", err, stderr.String())
	}

	if err != nil {
		return err
	}
	fmt.Printf("User %s created successfully\n", username)
	return nil
}

func (a *App) DeleteUser() error {

	distribution, err := getLinuxDistribution()

	var cmd *exec.Cmd
	if distribution == "ubuntu" || distribution == "debian" {
		cmd = exec.Command("deluser", "testuser")
	} else {
		cmd = exec.Command("userdel", "testuser")
	}

	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("User %s deleted successfully\n", "testuser")
	return nil
}

func (a *App) DeleteUserWithDir(username string) error {
	cmd := exec.Command("deluser", "--remove-home", username)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("User %s deleted successfully\n", username)
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

	// Get the Linux distribution
	distribution, err := getLinuxDistribution()

	var cmd *exec.Cmd

	if distribution == "ubuntu" || distribution == "debian" {
		cmd = exec.Command("apt", "list", "--installed")
	} else if distribution == "fedora" || distribution == "centos" || distribution == "rhel" {
		cmd = exec.Command("dnf", "list", "installed")
	} else if distribution == "arch" {
		cmd = exec.Command("pacman", "-Q")
	} else {
		fmt.Println("Unsupported Linux distribution")
		return nil
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Extract only the package names using the extractFirstParams function
	packageNames := ExtractFirstParams(string(out))
	return packageNames
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

func (a *App) GetCPUInfo() ([]cpu.InfoStat, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return cpuInfo, nil
}

///////////////////////////-----------------///////////////////////////

type CronJob struct {
	Schedule string
	Command  string
}

func (a *App) ListCronJobs() []CronJob {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(output), "\n")
	jobs := []CronJob{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 6 {
			continue
		}

		job := CronJob{
			Schedule: strings.Join(parts[:5], " "),
			Command:  strings.Join(parts[5:], " "),
		}

		jobs = append(jobs, job)
	}

	if jobs == nil || len(jobs) == 0 {
		fmt.Println("No cron jobs found")
	}

	return jobs
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
	// Create a temporary file to store the current crontab
	tmpfile, err := ioutil.TempFile("", "crontab")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	// Execute the crontab command with the '-l' option to list the current crontab
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// Write the current crontab to the temporary file
	_, err = tmpfile.Write(output)
	if err != nil {
		return err
	}

	// Add the new cron job to the temporary file
	_, err = fmt.Fprintf(tmpfile, "%s %s\n", schedule, command)
	if err != nil {
		return err
	}

	// Execute the crontab command with the temporary file as input to update the crontab
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
	}

	return listIptables()
}
