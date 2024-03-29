package backend

import (
	"fmt"
	"os/exec"
	"strconv"
)

func (b *Backend) CreateGroup(name string, gid *int) error {
	cmdArgs := []string{"groupadd", name}
	if gid != nil {
		cmdArgs = append(cmdArgs, "-g", strconv.Itoa(*gid))
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to create group: %s", name))
		return err
	}
	b.logger.Info(fmt.Sprintf("Group created successfully: %s", name))
	return nil
}

func (b *Backend) ModifyGroup(name string, gid int) error {
	cmd := exec.Command("groupmod", "-g", strconv.Itoa(gid), name)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to modify group: %s", name))
		return err
	}
	b.logger.Info(fmt.Sprintf("Group modified successfully: %s", name))
	return nil
}

func (b *Backend) DeleteGroup(name string) error {
	cmd := exec.Command("groupdel", name)
	err := cmd.Run()
	if err != nil {
		b.logger.Error(fmt.Sprintf("Failed to delete group: %s", name))
		return err
	}
	b.logger.Info(fmt.Sprintf("Group deleted successfully: %s", name))
	return nil
}
