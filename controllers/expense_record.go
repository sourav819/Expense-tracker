package controllers

import (
	"expense-tracker/constants"
	"expense-tracker/models"
	"expense-tracker/pkg/logger"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

var cache = &sync.Map{}

func (b *BaseController) AddExpense(c *gin.Context) {
	var (
		request       = ExpenseInfo{}
		err           error
		errMsg        = constants.ErrorEntity{}
		expenseRepo   = models.InitExpenseDetailsRepo(b.DB)
		expenseRecord = models.ExpenseDetails{}
		queryResp     = QueryResult{}
		targetRepo    = models.InitTargetDetails(b.DB)
		LimitExceeded *bool
		response      = ExpenseInfoResponse{}
		Newerr        error
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	value, ok := c.Get("user")
	if !ok {
		logger.Error("no value found for user")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg.GenerateError(http.StatusUnauthorized, "invalid user"))
		return
	}
	userDetails, ok := value.(models.User)
	if !ok {
		logger.Error("could not get detais for user")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	tx := b.DB.Begin()

	expenseDate, err := time.Parse("2006-01-02", request.DateOfExpense)
	if err != nil {
		logger.Error("invalid time")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid time, format should be YYYY-MM-DD"))
		return
	}

	startDate := strings.Split(now.BeginningOfMonth().String(), " ")
	endDate := strings.Split(now.EndOfMonth().String(), " ")
	cacheKey := startDate[0] + userDetails.UUID

	weekStart := strings.Split(now.BeginningOfWeek().String(), " ")
	weekEnd := strings.Split(now.EndOfWeek().String(), " ")

	allTargetInfo, err := targetRepo.GetAllDetailsOfTarget(tx, startDate[0], endDate[0], weekStart[0], weekEnd[0])
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("error in fetching target record from database")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	for _, value := range *allTargetInfo {
		//add current expense to all rows
		value.SpentAmount += request.Amount
		if value.SpentAmount > value.TargetAmount {
			//add expense
			*LimitExceeded = true
		}
	}

	//check for cache
	sumValues, ok := cache.Load(cacheKey)
	if !ok {
		query := fmt.Sprintf(`SELECT SUM(amount), count(amount) from %s `, ExpenseDetails)
		whereQuery := fmt.Sprintf(`WHERE date_of_expense >= '%s' AND date_of_expense <= '%s'`, startDate[0], endDate[0])
		err = b.DB.Raw(query + whereQuery).Scan(&queryResp).Error
		if err != nil {
			logger.Error("error in fetching record database")
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
		//put value to cache

		location, _ := time.LoadLocation("Asia/Kolkata")
		expenseDate = expenseDate.In(location)
		expenseRecord.Amount = request.Amount
		expenseRecord.Category = request.Category
		expenseRecord.DateOfExpense = &expenseDate
		expenseRecord.UserID = userDetails.ID
		expenseRecord.Remarks = request.Remarks
		if LimitExceeded != nil {
			expenseRecord.LimitExceed = true
		}

		err = expenseRepo.Create(tx, &expenseRecord)
		if err != nil {
			logger.Error("error in inserting record to database")
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
		//update spent amount record in target table
		for i := 0; i < len(*allTargetInfo); i++ {
			(*allTargetInfo)[i].SpentAmount += request.Amount
			Newerr = targetRepo.UpdateWithTx(tx, &(*allTargetInfo)[i])
			if Newerr != nil {
				break
			}
		}
		if Newerr != nil {
			logger.Error("error in updating record to database")
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}

		tx.Commit()
		queryResp.Sum += uint(request.Amount)
		queryResp.Count += 1
		cache.Store(startDate[0]+userDetails.UUID, queryResp)
		response.Message = "Data saved successfully"
		response.TotalExpenditure = request.Amount + uint64(queryResp.Sum)
		response.TotalNumberOfTransactions = queryResp.Count + 1
		if LimitExceeded != nil {
			response.Remarks = "you have exceeded your target!!"
		}

		c.JSON(http.StatusOK, response)
		return
	}

	location, _ := time.LoadLocation("Asia/Kolkata")
	expenseDate = expenseDate.In(location)
	expenseRecord.Amount = request.Amount
	expenseRecord.Category = request.Category
	expenseRecord.DateOfExpense = &expenseDate
	expenseRecord.UserID = userDetails.ID
	expenseRecord.Remarks = request.Remarks
	if LimitExceeded != nil {
		expenseRecord.LimitExceed = true
	}

	err = expenseRepo.Create(tx, &expenseRecord)
	if err != nil {
		logger.Error("error in inserting record to database")
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	for i := 0; i < len(*allTargetInfo); i++ {
		(*allTargetInfo)[i].SpentAmount += request.Amount
		Newerr = targetRepo.UpdateWithTx(tx, &(*allTargetInfo)[i])
		if Newerr != nil {
			break
		}
	}
	if Newerr != nil {
		logger.Error("error in updating record to database")
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	tx.Commit()

	logger.Info("fetching records from cache")
	expensesummary := sumValues.(QueryResult)
	expensesummary.Sum += uint(request.Amount)
	expensesummary.Count += 1
	cache.Store(cacheKey, expensesummary)

	c.JSON(http.StatusOK, gin.H{
		"message":                            "data saved successfully",
		"total expenditure in current month": expensesummary.Sum,
		"total number of transactions":       expensesummary.Count,
	})

}
