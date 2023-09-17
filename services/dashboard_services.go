package services

import (
	"expense-tracker/objects"
	"expense-tracker/pkg/logger"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

const (
	CurrentMonth = "current-month"
	CurrentWeek  = "current-week"
)

func SeparateFilterWise(filter string, DB *gorm.DB) (*objects.Dashboard, error) {
	var (
		query  = ""
		resp   = []objects.Response{}
		err    error
		total  uint
		result = objects.Dashboard{}
	)
	startDate, endDate := PrepareDates(filter)
	StartDate := strings.Split(*startDate, " ")
	EndDate := strings.Split(*endDate, " ")
	if startDate == nil && endDate == nil {
		now := time.Now()
		query += fmt.Sprintf(`SELECT SUM(amount), category FROM (select amount,category from expense_details where created_at='%s'`, now.Format("2006-01-02"))
	}

	query += fmt.Sprintf(`SELECT SUM(amount), category FROM (select amount,category from expense_details where created_at>='%s'
	AND created_at<='%s') as temp GROUP BY category`, StartDate[0], EndDate[0])

	err = DB.Raw(query).Scan(&resp).Error
	if err != nil {
		logger.Error("unable to dump response ", err)
		return nil, err
	}
	for i := 0; i < len(resp); i++ {
		total += resp[i].Sum
	}
	for i := 0; i < len(resp); i++ {
		perc := CalculateData(total, resp[i].Sum)
		resp[i].Percentage = perc

	}
	result.Data = append(result.Data, resp...)
	result.TotalExpenditure = fmt.Sprint(total)
	result.TotalTransactions = fmt.Sprint(len(resp))

	return &result, nil
}

func PrepareDates(filter string) (*string, *string) {
	var (
		startDate string
		endDate   string
	)
	if filter == CurrentMonth {
		startDate = now.BeginningOfMonth().String()
		endDate = now.EndOfMonth().String()
	} else if filter == CurrentWeek {
		startDate = now.BeginningOfWeek().String()
		endDate = now.EndOfWeek().String()
	} else {
		//today
		return nil, nil
	}
	return &startDate, &endDate
}

func CalculateData(total, sum uint) float64 {
	percentage := math.Max(float64(sum)/float64(total)*100, 0)
	if math.IsNaN(percentage) {
		return 0
	} else {
		percentage = math.Round(percentage*100) / 100
		return percentage
	}
}
