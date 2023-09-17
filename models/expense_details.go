package models

import (
	"time"

	"gorm.io/gorm"
)

type ExpenseDetails struct {
	ID            uint64         `gorm:"primarykey" json:"-"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	UserID        uint64         `json:"-"`
	User          User           `gorm:"constraint:OnDelete:CASCADE;"`
	Amount        uint64         `json:"amount"`
	DateOfExpense *time.Time     `json:"date_of_expense"`
	Category      string         `json:"category"`
	Remarks       string         `json:"remarks"`
	LimitExceed   bool           `json:"limit_exceed" gorm:"default:false"`
}

type ExpenseDetailsRepo struct {
	DB *gorm.DB
}

func (e *ExpenseDetailsRepo) Create(tx *gorm.DB, expense *ExpenseDetails) error {
	return e.CreatewithTx(tx, expense)
}

func (e *ExpenseDetailsRepo) CreatewithTx(tx *gorm.DB, expense *ExpenseDetails) error {

	err := tx.Create(&expense).Error
	if err != nil {
		return err
	}
	return nil
}
