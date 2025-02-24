package app

import (
	"log/slog"

	"github.com/dknathalage/pkg/log"
)

type App struct {
	Logger *slog.Logger
}

func NewApp() *App {
	return &App{
		Logger: log.NewJsonLogger(),
	}
}
