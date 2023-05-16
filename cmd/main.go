package main

import (
	"citizenship/internal/app"
	_ "golang.org/x/net/context"

	"citizenship/internal/config"
	"citizenship/pkg/logging"
	_ "context"
	_ "github.com/julienschmidt/httprouter"
)

func main() {
	// entry point
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info(cfg)
	// logger.Info("mongodb composite initializing")
	// mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, "", "", "", "")
	// if err != nil {
	// 	logger.Fatal("mongodb composite failed")
	// }

	app := application.NewApp()
	app.Run()
}

// Ctrl+L - select whole line
// Alt+J - find next such word with multiple cursor
// Ctrl+Shift+Arrow - move line up/down
// Ctrl+/ - (un)comment selection
// Ctrl+D - duplicate line
// Ctrl+Alt+L - format
