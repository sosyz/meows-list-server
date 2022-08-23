package utils

import "testing"

func TestErrorResponse(t *testing.T) {
	println(ErrorResponse("error"))
}

func TestSuccessResponse(t *testing.T) {
	println(SuccessResponse("success", "data"))
}

func TestMaskPhone(t *testing.T) {
	phone := MaskPhone("18888888888")
	println(phone)
}
