package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint64         `gorm:"primarykey" json:"-"`
	Created_at  *time.Time      `json:"created_at"`
	Updated_at  *time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UUID        string         `json:"uuid" gorm:"unique"`
	First_name  *string        `json:"first_name" validate:"required,min=2,max=100"`
	Last_name   *string        `json:"last_name" validate:"required,min=2,max=100"`
	Password    *string        `json:"Password" validate:"required,min=6"`
	Email       *string        `json:"email" validate:"email,required" gorm:"unique"`
	PhoneNumber *string        `json:"phone_number" validate:"required" gorm:"unique"`
	IsActive    *bool          `json:"is_active"`
}

type UserDetailsRepo struct {
	DB *gorm.DB
}

func (u *UserDetailsRepo) GetUserDetails(email, phoneNum string) (int64, error) {
	return u.GetwithTx(u.DB, &User{Email: &email, PhoneNumber: &phoneNum})
}

func (u *UserDetailsRepo) GetwithTx(tx *gorm.DB, email *User) (int64, error) {
	var ud User
	result := tx.Model(&User{}).Where(email).First(&ud)
	return result.RowsAffected, result.Error
}

// func (u *UserDetailsRepo) CheckUserBasedOnEmail(email string) (int64, error) {
// 	return u.GetwithTx(u.DB, &User{Email: &email})
// }

// func (u *UserDetailsRepo) CheckUserBasedOnPhone(phoneNum string) (int64, error) {
// 	return u.GetwithTx(u.DB, &User{PhoneNumber: &phoneNum})
// }

func (u *UserDetailsRepo) CheckDetailsForLogin(tx *gorm.DB, email, phoneNum string) (*User, error) {
	var us User
	err := tx.Model(&User{}).Where(`email=?`, email).Or(`phone_number=?`, phoneNum).First(&us).Error
	return &us, err
}
