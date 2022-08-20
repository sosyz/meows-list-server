package controller

import (
	"context"
	"encoding/json"
	"sonui.cn/meows-list-server/pkg/utils"
	"sonui.cn/meows-list-server/services"
)

type LoginParams struct {
	Email  string `form:"email" json:"name" valid:"Required; MaxSize(100); Email"`
	Pass   string `form:"password" json:"password" valid:"Required; MaxSize(100)"`
	Verify string `form:"verify" json:"verify"`
}

type SignupParams struct {
	Name  string `form:"name" json:"name" valid:"Required; MaxSize(100)"`
	Pass  string `form:"password" json:"password" valid:"Required; MaxSize(100)"`
	Email string `form:"email" json:"email" valid:"Required; MaxSize(100); Email"`
	Phone string `form:"phone" json:"phone"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

func UserLogin(ctx context.Context, opt *LoginParams) (string, error) {
	if next, msg := utils.Validation(opt); !next {
		res, _ := json.Marshal(utils.ErrorResponse(msg))
		return string(res), nil
	} else {
		if err := services.UserLogin(opt.Email, opt.Pass); err != nil {
			res, _ := json.Marshal(utils.ErrorResponse(err.Error()))
			return string(res), nil
		}
		res, _ := json.Marshal(utils.SuccessResponse("登录成功", LoginResponse{
			Token: "",
		}))
		return string(res), nil
	}
}

func UserSignup(ctx context.Context, opt *SignupParams) (string, error) {
	if next, msg := utils.Validation(opt); !next {
		res, _ := json.Marshal(utils.ErrorResponse(msg))
		return string(res), nil
	} else {
		if err := services.UserSignup(opt.Name, opt.Pass, opt.Email, opt.Phone); err != nil {
			res, _ := json.Marshal(utils.ErrorResponse(err.Error()))
			return string(res), nil
		}
		res, _ := json.Marshal(utils.SuccessResponse("登录成功", SignupResponse{
			Message: "注册成功",
		}))
		return string(res), nil
	}
}
