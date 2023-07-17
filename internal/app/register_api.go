package app

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/fairytale5571/ipcurrency/docs"
	"github.com/fairytale5571/ipcurrency/internal/api/delivery/http/ipinfo"
	ipInfoRepository "github.com/fairytale5571/ipcurrency/internal/api/repository/ipinfo"
	ipInfoService "github.com/fairytale5571/ipcurrency/internal/api/services/ipinfo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *App) registerRepositories() {
	app.ipInfoRepo = ipInfoRepository.NewRepository()
}

func (app *App) registerServices() {
	app.ipInfoService = ipInfoService.NewService(app.ipInfoRepo)
}

func (app *App) registerHandlers() {
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
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/ip-info", app.ipInfoHTTPHandler.GetIPInfo())
}
