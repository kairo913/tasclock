package main

import (
	"context"
	"embed"
	"log"

	"github.com/kairo913/tasclock/todo"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	todo := todo.NewTodo()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "TasClock",
		Width:  1024,
		MinWidth: 1024,
		Height: 768,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			todo.Startup()
		},
		OnShutdown: func(ctx context.Context) {
			todo.Shutdown()
		},
		Bind: []interface{}{
			app,
			todo,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
