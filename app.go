package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	cmd := exec.Command("adduser", "testuser")
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
	cmd := exec.Command("deluser", "testuser")
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