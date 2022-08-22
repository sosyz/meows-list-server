package controller

import (
	"context"
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

func UserLogin(ctx context.Context, opt *LoginParams) string {
	if next, msg := utils.Validation(opt); !next {
		return utils.ErrorResponse(msg)
	} else {
		if token, err := services.UserLogin(opt.Email, opt.Pass); err != nil {
			return utils.ErrorResponse(err.Error())
		} else {
			return utils.SuccessResponse("登录成功", LoginResponse{
				Token: token,
			})
		}
	}
}

func UserSignup(ctx context.Context, opt *SignupParams) string {
	if next, msg := utils.Validation(opt); !next {
		return utils.ErrorResponse(msg)
	} else {
		if err := services.UserSignup(opt.Name, opt.Pass, opt.Email, opt.Phone); err != nil {
			return utils.ErrorResponse(err.Error())
		}
		return utils.SuccessResponse("登录成功", SignupResponse{
			Message: "注册成功",
		})
	}
}
