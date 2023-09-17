package controllers

import (
	"expense-tracker/constants"
	"expense-tracker/pkg/logger"
	"expense-tracker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BaseController) Dashboard(c *gin.Context) {
	var (
		err    error
		errMsg = constants.ErrorEntity{}
	)
	filter, ok := c.GetQuery("filter")
	if !ok {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request,pass filter query"))
		return
	}
	resp, err := services.SeparateFilterWise(filter, b.DB)
	if err != nil {
		logger.Error("unable to fetch details for user")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	c.JSON(http.StatusOK, *resp)

}
