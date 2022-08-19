package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/validation"
	"sonui.cn/meows-list-server/pkg/utils"
	"sonui.cn/meows-list-server/services"
)

type LoginParams struct {
	Name   string `form:"name" json:"name" valid:"Required; MaxSize(100)"`
	Pass   string `form:"password" json:"password" valid:"Required; MaxSize(100)"`
	Verify string `form:"verify" json:"verify"`
}

func (l *LoginParams) Valid(v *validation.Validation) {
}

func UserLogin(ctx context.Context, opt *LoginParams) (string, error) {
	var (
		ret string
		err error
	)
	// 参数检查
	if next, msg := utils.Validation(opt); !next {
		err = errors.New(msg)
	} else {
		var res interface{}
		res, err = services.UserLogin(opt.Name, opt.Pass)
		_json, _ := json.Marshal(res)
		ret = string(_json)
	}
	return ret, err
}
