package controllers

import (
	"expense-tracker/constants"
	"expense-tracker/models"
	"expense-tracker/pkg/logger"
	"expense-tracker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (b *BaseController) SignUpUser(c *gin.Context) {

	var (
		request     = SignUpRequest{}
		err         error
		errMsg      = constants.ErrorEntity{}
		userRepo    = models.InitUserDetailsRepo(b.DB)
		userDetails = models.User{}
	)

	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	tx := b.DB.Begin()
	recordCount, err := userRepo.GetUserDetails(request.Email, request.PhoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("unable to get user details")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	if recordCount != 0 {
		logger.Error("user already exists")
		c.AbortWithStatusJSON(http.StatusForbidden, errMsg.GenerateError(http.StatusForbidden, "user already exists"))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("unable to hash password")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	randomId, err := utils.CreateNanoID(20)
	if err != nil {
		logger.Error("unable to generate random uuid")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	UserIsActive := true
	passwd := string(hashedPassword)
	userDetails.First_name = &request.FirstName
	userDetails.Last_name = &request.LastName
	userDetails.Email = &request.Email
	userDetails.PhoneNumber = &request.PhoneNumber
	userDetails.Password = &passwd
	userDetails.IsActive = &UserIsActive
	userDetails.UUID = "usr_" + randomId

	err = tx.Create(&userDetails).Error
	if err != nil {
		logger.Error("unable to insert record of the user")
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))

	}
	err = tx.Commit().Error
	if err != nil {
		b.Log.Error("unable to commit the code and url ", err)
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message ": "user signed up successfully"})
}
