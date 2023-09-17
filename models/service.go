package models

import "gorm.io/gorm"

func InitUserDetailsRepo(DB *gorm.DB) IUserDetails {
	return &UserDetailsRepo{
		DB: DB,
	}
}

func InitExpenseDetailsRepo(DB *gorm.DB) IExpenseDetails {
	return &ExpenseDetailsRepo{
		DB: DB,
	}
}

func InitTargetDetails(DB *gorm.DB) ITargetDetails {
	return &TargetDetailsRepo{
		DB: DB,
	}
}
