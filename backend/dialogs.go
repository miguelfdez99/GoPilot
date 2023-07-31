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

func (b *Backend) OpenDialogInfo(msg string) {
	runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Info",
		Message: msg,
	})
}

func (b *Backend) OpenDialogError(msg string) {
	runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Error",
		Message: msg,
	})
}

func (b *Backend) OpenDialogWarning(msg string) {
	runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.WarningDialog,
		Title:   "Warning",
		Message: msg,
	})
}

func (b *Backend) OpenDialogQuestion(msg string) string {
	result, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Question",
		Message:       msg,
		DefaultButton: "No",
	})
	if err != nil {
		return ""
	}

	return result
}

func (b *Backend) OpenURL(url string) {
	runtime.BrowserOpenURL(b.ctx, url)
}
