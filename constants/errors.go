package constants

type ErrorEntity struct {
	StatusCode       int    `json:"status_code"`
	ErrorDescription string `json:"error_description"`
	Message          string `json:"message"`
}

func (e *ErrorEntity) GenerateError(statuscode int, message string) *ErrorEntity {
	e.StatusCode = statuscode
	e.ErrorDescription = GetErrorMessage(statuscode)
	e.Message = message

	return e
}


