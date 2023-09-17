package controllers

import (
	"expense-tracker/pkg/config"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config config.Config
}
