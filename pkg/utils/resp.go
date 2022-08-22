package utils

import "encoding/json"

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

func ErrorResponse(message string) string {
	ret := err{}
	ret.Success = false
	ret.Message = message
	return packResponse(ret)
}

func SuccessResponse(message string, data interface{}) string {
	ret := success{}
	ret.Success = true
	ret.Message = message
	ret.Data = data
	return packResponse(ret)
}

func packResponse(data interface{}) string {
	res, _ := json.Marshal(data)
	return string(res)
}
