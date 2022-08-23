package utils

import (
	"github.com/beego/beego/v2/core/validation"
)

func Validation(data interface{}) (bool, string) {
	validate := &validation.Validation{}
	b, err := validate.Valid(data)
	if err != nil {
		return false, err.Error()
	}
	if !b {
		for _, err := range validate.Errors {
			return false, err.Key + " " + err.Message
		}
	}
	return true, ""
}
