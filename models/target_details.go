package models

import (
	"time"

	"gorm.io/gorm"
)

type TargetDetails struct {
	ID           uint64         `gorm:"primarykey" json:"-"`
	Created_at   *time.Time     `json:"created_at"`
	Updated_at   *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	UserID       uint64         `json:"-"`
	User         User           `gorm:"constraint:OnDelete:CASCADE;"`
	TargetType   string         `json:"target_type"` //filter
	TargetAmount uint64         `json:"target_amount"`
	SpentAmount  uint64         `json:"spent_amount"`
	Startdate    string         `json:"start_date"`
	EndDate      string         `json:"end_date"`
	Suggestions  string         `json:"suggestions"`
}

type TargetDetailsRepo struct {
	DB *gorm.DB
}

func (t *TargetDetailsRepo) GetTargetDetails(startDate, endDate string) (*TargetDetails, error) {
	return t.GetwithTx(t.DB, &TargetDetails{Startdate: startDate, EndDate: endDate})
}

func (t *TargetDetailsRepo) GetwithTx(tx *gorm.DB, dateRange *TargetDetails) (*TargetDetails, error) {
	var td TargetDetails
	err := tx.Model(&TargetDetails{}).Where(dateRange).Find(&td).Error
	if err != nil {
		return nil, err
	}
	return &td, nil
}

func (t *TargetDetailsRepo) UpdateWithTx(tx *gorm.DB, updateValues *TargetDetails) error {
	err := tx.Model(&TargetDetails{}).Where(&TargetDetails{ID: updateValues.ID}).Updates(updateValues).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TargetDetailsRepo) GetAllDetailsOfTarget(tx *gorm.DB, startdateMonth, enddateMonth, weekStart, weekEnd string) (*[]TargetDetails, error) {
	var td []TargetDetails
	err := tx.Model(&TargetDetails{}).Where(&TargetDetails{Startdate: startdateMonth, EndDate: enddateMonth}).
		Or(&TargetDetails{Startdate: weekStart, EndDate: weekEnd}).Find(&td).Error
	if err != nil {
		return nil, err
	}
	return &td, nil
}
