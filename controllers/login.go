package controllers

import (
	"expense-tracker/auth"
	"expense-tracker/constants"
	"expense-tracker/models"
	"expense-tracker/pkg/logger"
	"expense-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (b *BaseController) Login(c *gin.Context) {
	var (
		request  = LoginRequest{}
		err      error
		errMsg   = constants.ErrorEntity{}
		userRepo = models.InitUserDetailsRepo(b.DB)
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	tx := b.DB.Begin()
	//checking if the phone number or email exists
	usrDetails, err := userRepo.CheckDetailsForLogin(tx, request.Email, request.PhoneNumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Info("user doesnot exists")
			c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "emailId or phone number does not exists, please signup"))
			return
		} else {
			logger.Error("unable to fetch details for user")
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
	}
	//check password
	if !utils.VerifyPassword(*usrDetails.Password, request.Password) {
		logger.Error("password doesnot match")
		c.AbortWithStatusJSON(http.StatusForbidden, errMsg.GenerateError(http.StatusForbidden, "incorrect password"))
		return
	}
	token, err := auth.GenerateToken(b.Config.JWTConfig.JWTTokenValidTimeInHour, *usrDetails.Email, *usrDetails.First_name, *usrDetails.Last_name, usrDetails.UUID, b.Config.JWTConfig.JWTSecret)
	if err != nil {
		logger.Error("unable to generate token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})

}
