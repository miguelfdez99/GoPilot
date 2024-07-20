package backend

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (b *Backend) CheckAdmin() bool {
	uid := os.Getuid()
	return uid == 0
}

func (b *Backend) CreateUser(user User) error {
	args := []string{
		"-m",
		"-s", user.Shell,
		"-p", user.Password,
		"-g", fmt.Sprint(user.GID),
	}

	if user.UID != 0 {
		args = append(args, "-u", fmt.Sprint(user.UID))
	}

	args = append(args, user.Username)

	cmd := exec.Command("useradd", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to create user: %s. Reason: %s", user.Username, stderr.String()))
		return fmt.Errorf("failed to create user: %v. Reason: %s", err, stderr.String())
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
