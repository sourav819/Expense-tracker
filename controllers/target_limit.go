package controllers

import (
	"expense-tracker/constants"
	"expense-tracker/models"
	"expense-tracker/pkg/logger"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

func (b *BaseController) SetTarget(c *gin.Context) {
	var (
		err           error
		errMsg        = constants.ErrorEntity{}
		request       = TargetRequest{}
		targetDetails = models.TargetDetails{}
		startdate     string
		enddate       string
		targetRepo    = models.InitTargetDetails(b.DB)
		response      = TargetResponse{}
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	value, ok := c.Get("user")
	if !ok {
		logger.Error("no value found for user ", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg.GenerateError(http.StatusUnauthorized, "invalid user"))
		return
	}
	userDetails, ok := value.(models.User)
	if !ok {
		logger.Error("could not get detais for user ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	tx := b.DB.Begin()
	//prepare range

	if strings.ToLower(request.Filter) == Monthly {
		startdate = now.BeginningOfMonth().String()
		enddate = now.EndOfMonth().String()
		startTime := strings.Split(startdate, " ")
		endTime := strings.Split(enddate, " ")
		startdate = startTime[0]
		enddate = endTime[0]

	} else {
		startdate = now.BeginningOfWeek().String()
		enddate = now.EndOfWeek().String()
		startTime := strings.Split(startdate, " ")
		endTime := strings.Split(enddate, " ")
		startdate = startTime[0]
		enddate = endTime[0]
	}

	targetDetails.UserID = userDetails.ID
	targetDetails.TargetType = request.Filter
	targetDetails.TargetAmount = request.Amount
	targetDetails.SpentAmount = 0
	targetDetails.Startdate = startdate
	targetDetails.EndDate = enddate
	//check if a target already exists
	targDetails, err := targetRepo.GetTargetDetails(startdate, enddate)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("unable to fetch target details ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	} else if err == gorm.ErrRecordNotFound {
		err = tx.Create(&targetDetails).Error
		if err != nil {
			tx.Rollback()
			logger.Error("unable to create target ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
		response.Message = "Target created successfully, All the best!"
		response.AmountSet = request.Amount
		response.TimeRange = fmt.Sprintf("from %s to %s ", startdate, enddate)
	} else {
		// update target
		valuesToUpdate := models.TargetDetails{}
		valuesToUpdate.ID = targDetails.ID
		valuesToUpdate.TargetAmount = request.Amount
		err = targetRepo.UpdateWithTx(b.DB, &valuesToUpdate)
		if err != nil {
			logger.Error("unable to update target details ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
		response.Message = "Target updated successfully, All the best!"
		response.AmountSet = request.Amount
		response.TimeRange = fmt.Sprintf("from %s to %s ", startdate, enddate)
	}

	tx.Commit()
	c.JSON(http.StatusOK, response)
}
