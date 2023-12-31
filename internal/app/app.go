package app

import (
	"net/http"

	"github.com/fairytale5571/ipcurrency/internal/api/delivery"
	"github.com/fairytale5571/ipcurrency/internal/api/repository"
	"github.com/fairytale5571/ipcurrency/internal/api/services"
)

type App struct {
	server *http.Server

	ipInfoService     services.IPInfo
	ipInfoRepo        repository.IPInfo
	ipInfoHTTPHandler delivery.IPInfoHTTP
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {
	app.registerRepositories()
	app.registerServices()
	app.registerHandlers()
	return app.setupServerAndRoutes()
}

func (app *App) Stop() error {
	return app.server.Close()
}
