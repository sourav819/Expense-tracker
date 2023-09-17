package constants

const (
	InternalServerError = 500
	BadRequest          = 400
	Forbidden           = 403
	ServiceUnavailable  = 503
	Unauthorized        = 401
)

var StatusMapping = map[int]string{
	InternalServerError: "Internal Server Error",
	BadRequest:          "Bad Request",
	Forbidden:           "Forbidden Error",
	ServiceUnavailable:  "Service currently unavilable",
	Unauthorized:        "Unauthorized error",
}

func GetErrorMessage(code int) string {
	return StatusMapping[code]
}
