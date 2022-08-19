package utils

type err struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type success struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(message string) interface{} {
	return err{
		Success: false,
		Message: message,
	}
}

func SuccessResponse(message string, data interface{}) interface{} {
	return success{
		Success: true,
		Message: message,
		Data:    data,
	}
}
