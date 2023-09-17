package objects

type (
	Dashboard struct {
		TotalExpenditure  string     `json:"total_expenditure"`
		TotalTransactions string     `json:"total_transactions"`
		Data              []Response `json:"data"`
	}
	Response struct {
		Sum        uint    `json:"sum"`
		Category   string  `json:"category"`
		Percentage float64 `json:"percentage"`
	}
)
