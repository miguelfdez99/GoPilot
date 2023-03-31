package back

import (
	"context"
	"fmt"
	"os/exec"
)

type Backend struct {
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

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) CreateUser2(user User) error {
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
