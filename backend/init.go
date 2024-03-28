package backend

import (
	"context"
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
