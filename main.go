package main

import (
	"expense-tracker/pkg/config"
	"expense-tracker/pkg/database"
	"expense-tracker/pkg/logger"
	"expense-tracker/routers"
)

func main() {
	//reading config from env
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("unable to load config ", err)
	}
	//making the database connection
	db, err := database.GetDB(cfg)
	if err != nil {
		logger.Fatal("unable to load config ", err)
	}
	app := config.AppConfig{
		DB:     db,
		Config: cfg,
	}

	routers.SetupAndRunServer(&app)
}
