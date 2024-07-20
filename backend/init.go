package backend

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
)

type Backend struct {
	ctx       context.Context
	logger    *CustomLogger
	watchList sync.Map
}

func NewBackend(l *CustomLogger) *Backend {
	return &Backend{
		logger: l,
	}
}

func (b *Backend) Startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *Backend) CommandExists(cmd string) (bool, string) {
	_, err := exec.LookPath(cmd)
	if err != nil {
		b.logger.Error(fmt.Sprintf("%s command is not available. Error: %v", cmd, err))
		return false, fmt.Sprintf("%s command is not available", cmd)
	}
	return true, ""
}
