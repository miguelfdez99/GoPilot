package backend

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type Backend struct {
	ctx    context.Context
	logger logger.Logger
}

func NewBackend(l logger.Logger) *Backend {
	return &Backend{
		logger: l,
	}
}

func (b *Backend) startup(ctx context.Context) {
	b.ctx = ctx
}
