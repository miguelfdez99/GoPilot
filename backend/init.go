package backend

import "context"

// App struct
type Backend struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewBackend() *Backend {
	return &Backend{}
}

// startup is called when the Backend starts. The context is saved
// so we can call the runtime methods
func (b *Backend) startup(ctx context.Context) {
	b.ctx = ctx
}
