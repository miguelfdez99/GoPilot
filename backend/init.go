package backend

import (
	"context"
	"sync"

	"github.com/docker/docker/client"
)

type Backend struct {
	ctx       context.Context
	logger    *CustomLogger
	watchList sync.Map
	cli       *client.Client
}

func NewBackend(l *CustomLogger) (*Backend, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &Backend{
		logger: l,
		cli:    cli,
	}, nil
}

func (b *Backend) Startup(ctx context.Context) {
	b.ctx = ctx
}
