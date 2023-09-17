package controllers

//authentication structs
type (
	SignUpRequest struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}

	LoginRequest struct {
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Password    string `json:"password"`
	}
)

type (
	ExpenseInfo struct {
		Amount        uint64 `json:"amount"`
		Category      string `json:"category"`
		DateOfExpense string `json:"date_of_expense"` //YYYY-MM-DD
		Remarks       string `json:"remarks,omitempty"`
	}
	QueryResult struct {
		Sum   uint `json:"sum"`
		Count uint `json:"count"`
	}
	ExpenseInfoResponse struct {
		Message                   string `json:"message"`
		TotalExpenditure          uint64 `json:"total_expenditure"`
		TotalNumberOfTransactions uint   `json:"total_number_of_transactions"`
		Remarks                   string `json:"remarks"`
	}
)

type (
	TargetRequest struct {
		Filter string `json:"filter"`
		Amount uint64 `json:"amount"`
	}
	TargetResponse struct {
		Message   string `json:"message"`
		TimeRange string `json:"time_range"`
		AmountSet uint64 `json:"amount_set"`
	}
)

const (
	ExpenseDetails = "expense_details"
	Monthly        = "monthly"
	Weekly         = "weekly"
)
