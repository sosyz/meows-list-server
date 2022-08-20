package utils

type all struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type err struct {
	all
}

type success struct {
	all
	Data interface{} `json:"data"`
}

func ErrorResponse(message string) interface{} {
	ret := err{}
	ret.Success = false
	ret.Message = message
	return ret
}

func SuccessResponse(message string, data interface{}) interface{} {
	ret := success{}
	ret.Success = true
	ret.Message = message
	ret.Data = data
	return ret
}
