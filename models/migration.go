package models

var migrationModels = []interface{}{
	//table name here
	&User{},
	&ExpenseDetails{},
	&TargetDetails{},
}

func GetMigrationModels() []interface{} {
	return migrationModels
}
