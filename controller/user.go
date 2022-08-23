package controller

import (
	"context"
	"sonui.cn/meows-list-server/models"
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

func UserRegister(ctx context.Context, opt *SignupParams) string {
	if next, msg := utils.Validation(opt); !next {
		return utils.ErrorResponse(msg)
	} else {
		if err := services.UserRegister(opt.Name, opt.Pass, opt.Email, opt.Phone); err != nil {
			return utils.ErrorResponse(err.Error())
		}
		return utils.SuccessResponse("登录成功", SignupResponse{
			Message: "注册成功",
		})
	}
}

func UserInfo(ctx context.Context) string {
	token := ctx.Value("token").(string)
	if v := services.GetUserByToken(token); v == nil {
		return utils.ErrorResponse("未登录")
	} else {
		// 去除敏感信息
		user := models.User{
			Name:  v.Name,
			Email: v.Email,
			Phone: v.Phone,
		}
		if user.Phone != "" {
			user.Phone = utils.MaskPhone(user.Phone)
		}
		return utils.SuccessResponse("获取用户信息成功", user)
	}
}
