package routers

import (
	"net/http"
	"expense-tracker/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupAndRunServer(app *config.AppConfig) {
	environment := app.Config.Server.Debug
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	if app.Config.Server.VPCProxyCIDR != "" {
		router.SetTrustedProxies([]string{app.Config.Server.VPCProxyCIDR})
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middleware.CORSMiddleware())
	app.Router = router

	// Register routes
	registerRoutes(*app)
	// Run Server after InitRoutes
	runServer(*app)
}

func registerRoutes(app config.AppConfig) {
	app.Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	// Register All routes
	v1Routes(app)
}
