package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

// App struct
type App struct {
	ctx context.Context
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

	distribution := checkDistro()

	var cmd *exec.Cmd
	if distribution == "ubuntu" || distribution == "debian" {
		cmd = exec.Command("adduser", "testuser")
	} else {
		cmd = exec.Command("useradd", "testuser")
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if stderr.Len() > 0 && stderr.String() == "adduser: The user `testuser' already exists.\n" {
		fmt.Printf("error creating user: %v\nStderr output: %s", err, stderr.String())
	}

	if err != nil {
		return err
	}
	fmt.Printf("User %s created successfully\n", "testuser")
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

	distribution := checkDistro()

	var cmd *exec.Cmd
	if distribution == "ubuntu" || distribution == "debian" {
		cmd = exec.Command("deluser", "testuser")
	} else {
		cmd = exec.Command("userdel", "testuser")
	}

	err := cmd.Run()
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

func (a *App) ListPackages() []string {

	// Get the Linux distribution
	distribution := checkDistro()

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
	s := packageNames[1:25]
	fmt.Println(s)
	return s
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

func (a *App) GetCPUInfo() ([]cpu.InfoStat, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return cpuInfo, nil
}
