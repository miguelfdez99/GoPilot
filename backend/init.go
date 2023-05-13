package backend

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type Backend struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewBackend() *Backend {
	logger := logrus.New()
	logger.Out = os.Stdout

	// Set log level, e.g. logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel
	logger.SetLevel(logrus.InfoLevel)

	// Use JSONFormatter for structured logging
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Backend{
		logger: logger,
	}
}

func (b *Backend) startup(ctx context.Context) {
	b.ctx = ctx
}
