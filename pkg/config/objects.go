package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppConfig struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config Config
}

type Config struct {
	Server    ServerConfiguration   `mapstructure:",squash"`
	Database  DatabaseConfiguration `mapstructure:",squash"`
	JWTConfig JWTConfig             `mapstructure:",squash"`
}

type ServerConfiguration struct {
	Debug                bool   `mapstructure:"DEBUG"`
	Port                 string `mapstructure:"SERVER_PORT"`
	VPCProxyCIDR         string `mapstructure:"VPC_CIDR"`
	LimitCountPerRequest int64
}

type DatabaseConfiguration struct {
	Name     string `mapstructure:"MAIN_DB_NAME"`
	Username string `mapstructure:"MAIN_DB_USER"`
	Password string `mapstructure:"MAIN_DB_PASSWORD"`
	Host     string `mapstructure:"MAIN_DB_HOST"`
	Port     string `mapstructure:"MAIN_DB_PORT"`
	SSLMode  string `mapstructure:"MAIN_DB_SSL_MODE"`
}

type JWTConfig struct {
	JWTSecret               string `mapstructure:"JWT_SECRET"`
	JWTTokenValidTimeInHour uint32 `mapstructure:"JWT_TOKEN_VALID_TIME_IN_HOUR"`
}
