package routers

import (
	"expense-tracker/controllers"
	"expense-tracker/pkg/config"
	"expense-tracker/routers/middleware"

	"github.com/sirupsen/logrus"
)

func v1Routes(app config.AppConfig) {
	ctrl := controllers.BaseController{
		Config: app.Config,
		DB:     app.DB,
		Log:    logrus.New(),
	}
	v1 := app.Router.Group("/v1")

	//urls entity
	v1.POST("/sign-up", ctrl.SignUpUser)
	v1.POST("/login", ctrl.Login)
	authenticateGroup := v1.Group("/authenticate")
	authenticateGroup.Use(middleware.Authentication(ctrl.Config.JWTConfig.JWTSecret, ctrl.DB))
	authenticateGroup.POST("/check", ctrl.Check)
	authenticateGroup.POST("/addExpense", ctrl.AddExpense)
	authenticateGroup.GET("/dashboard", ctrl.Dashboard)
	authenticateGroup.POST("/setTarget",ctrl.SetTarget)

}
