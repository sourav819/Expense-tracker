package models

import "gorm.io/gorm"

type IUserDetails interface {
	GetUserDetails(email, phoneNum string) (int64, error)
	GetwithTx(tx *gorm.DB, email *User) (int64, error)
	CheckDetailsForLogin(tx *gorm.DB, email, phoneNum string) (*User, error)
}

type IExpenseDetails interface {
	Create(tx *gorm.DB, expense *ExpenseDetails) error
	CreatewithTx(tx *gorm.DB, expense *ExpenseDetails) error
}

type ITargetDetails interface{
	GetTargetDetails(startDate, endDate string) (*TargetDetails, error)
	GetwithTx(tx *gorm.DB, dateRange *TargetDetails) (*TargetDetails, error)
	UpdateWithTx(tx *gorm.DB, updateValues *TargetDetails) error
	GetAllDetailsOfTarget(tx *gorm.DB, startdateMonth, enddateMonth, weekStart, weekEnd string) (*[]TargetDetails, error)
}