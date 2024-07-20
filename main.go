package main

import (
	"context"
	"embed"
	"goPilot/backend"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed assets/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create backend
	logFilePath := filepath.Join("logs.txt")
	customLogger, err := backend.NewCustomLogger(logger.NewDefaultLogger(), logFilePath)
	if err != nil {
		println("Error creating custom logger:", err.Error())
		return
	}
	defer customLogger.Close()
	backend := backend.NewBackend(customLogger)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "GoPilot",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			backend.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			backend,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
