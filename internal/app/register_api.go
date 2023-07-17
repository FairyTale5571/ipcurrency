package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fairytale5571/ipcurrency/internal/api/delivery/http/ipinfo"
	"github.com/fairytale5571/ipcurrency/internal/api/delivery/http/status"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	ipInfoRepository "github.com/fairytale5571/ipcurrency/internal/api/repository/ipinfo"
	ipInfoService "github.com/fairytale5571/ipcurrency/internal/api/services/ipinfo"
)

func (app *App) registerRepositories() {
	app.ipInfoRepo = ipInfoRepository.NewRepository()
}

func (app *App) registerServices() {
	app.ipInfoService = ipInfoService.NewService(app.ipInfoRepo)
}

func (app *App) registerHandlers() {
	app.statusHTTPHandler = status.NewHandler()
	app.ipInfoHTTPHandler = ipinfo.NewHandler(time.Now, app.ipInfoService)
}

func (app *App) setupServerAndRoutes() error {
	router := gin.Default()
	app.registerRoutes(router)

	app.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: router,
	}

	return app.server.ListenAndServe()
}

func (app *App) registerRoutes(router *gin.Engine) {
	router.GET("/status", app.statusHTTPHandler.CheckStatus())
	router.POST("/ip-info", app.ipInfoHTTPHandler.GetIPInfo())
}
