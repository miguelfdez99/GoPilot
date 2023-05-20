package backend

import (
	"fmt"
	"os/exec"
	"strings"
)

type User struct {
	Username      string
	Password      string
	UID           int
	GID           int
	HomeDirectory string
	Shell         string
	Groups        []string
}

func (b *Backend) CreateUser(user User) error {
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
		b.logger.Error(fmt.Sprintf("Failed to create user: %s", user.Username))
		return fmt.Errorf("failed to create user: %v", err)
	}
	b.logger.Info(fmt.Sprintf("User created successfully: %s", user.Username))
	return nil
}

func (b *Backend) DeleteUser(username string, removeHomeDir bool, forceDelete bool) error {
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
		b.logger.Error(fmt.Sprintf("Failed to delete user: %s", username))
		return fmt.Errorf("error deleting user %s: %s", username, err)
	}
	b.logger.Info(fmt.Sprintf("User deleted successfully: %s", username))
	return nil
}

func (b *Backend) ModifyUser(username string, userPtr *User) error {
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
		b.logger.Error(fmt.Sprintf("Failed to modify user: %s", username))
		return fmt.Errorf("failed to modify user: %v", err)
	}
	b.logger.Info(fmt.Sprintf("User modified successfully: %s", username))
	return nil
}
