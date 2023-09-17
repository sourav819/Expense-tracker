package controllers

import (
	// "expense-tracker/auth"
	"expense-tracker/models"
	"expense-tracker/pkg/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BaseController) Check(c *gin.Context) {
	value, ok := c.Get("user")
	if !ok {
		logger.Error("no value found")
	}
	fmt.Println("printing value ", value)
	details, ok := value.(models.User)
	if !ok {
		logger.Error("no value found 2")
	}

	c.JSON(http.StatusOK, gin.H{
		"firstname": *details.First_name,
		"lastname":*details.Last_name,
		"email":*details.Email,
		"phonenumber":*details.PhoneNumber,
		"id":details.ID,
	})
}
