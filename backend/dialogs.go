package backend

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (b *Backend) OpenFile() (string, error) {
	selection, err := runtime.OpenFileDialog(b.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}

func (b *Backend) OpenDir() (string, error) {
	selection, err := runtime.OpenDirectoryDialog(b.ctx, runtime.OpenDialogOptions{
		Title: "Select Directory",
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}
