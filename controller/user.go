package controller

import (
	"context"
	"github.com/beego/beego/v2/core/validation"
	"sonui.cn/meows-list-server/pkg/utils"
	"sonui.cn/meows-list-server/services"
)

type LoginParams struct {
	Email  string `form:"email" json:"name" valid:"Required; MaxSize(100); Email"`
	Pass   string `form:"password" json:"password" valid:"Required; MaxSize(100)"`
	Verify string `form:"verify" json:"verify"`
}

type RegisterParams struct {
	Name  string `form:"name" json:"name" valid:"Required; MaxSize(20)"`
	Pass  string `form:"password" json:"password" valid:"Required; MaxSize(32); MinSize(6)"`
	Email string `form:"email" json:"email" valid:"Required; MaxSize(100); Email"`
	Phone string `form:"phone" json:"phone"`
}

type UpdateParams struct {
	Name    string `form:"name" json:"name"`
	Email   string `form:"email" json:"email"`
	Phone   string `form:"phone" json:"phone"`
	OldPass string `form:"old_password" json:"old_password"`
	Pass    string `form:"password" json:"password"`
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

func UserRegister(ctx context.Context, opt *RegisterParams) string {
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
	// 获取请求头中的token
	token := ctx.Value("token").(string)
	if v := services.GetUserByToken(token); v == nil {
		return utils.ErrorResponse("未登录")
	} else {
		// 去除敏感信息
		var user = map[string]string{
			"name":  v.Name,
			"email": v.Email,
			"phone": v.Phone,
		}
		if user["phone"] != "" {
			user["phone"] = utils.MaskPhone(user["phone"])
		}
		return utils.SuccessResponse("获取用户信息成功", user)
	}
}

func UserUpdate(ctx context.Context, opt *UpdateParams) string {
	// 获取请求头中的token
	token := ctx.Value("token").(string)

	valid := validation.Validation{}
	if opt.Name != "" {
		valid.MaxSize(opt.Name, 20, "name").Message("名称最大长度为20")
	}
	if opt.Email != "" {
		valid.Email(opt.Email, "email").Message("邮箱格式不正确")
	}
	if opt.Phone != "" {
		valid.Mobile(opt.Phone, "phone").Message("手机号格式不正确")
	}
	if opt.Pass != "" || opt.OldPass != "" {
		valid.Required(opt.OldPass, "old_password").Message("旧密码不能为空")
		valid.Required(opt.Pass, "password").Message("新密码不能为空")
		valid.MaxSize(opt.Pass, 32, "password").Message("新密码最大长度为32")
		valid.MinSize(opt.Pass, 6, "password").Message("新密码最小长度为6")
		valid.MaxSize(opt.OldPass, 32, "password").Message("旧密码最大长度为32")
		valid.MinSize(opt.OldPass, 6, "password").Message("旧密码最小长度为6")
	}

	if valid.HasErrors() {
		return utils.ErrorResponse(valid.Errors[0].Message)
	} else {
		user := services.GetUserByToken(token)
		if user == nil {
			return utils.ErrorResponse("未登录")
		} else {
			if err := services.UserUpdate(token, opt.Name, opt.Email, opt.Phone, opt.OldPass, opt.Pass); err != nil {
				return utils.ErrorResponse(err.Error())
			} else {
				return utils.SuccessResponse("更新成功", "success")
			}
		}
	}
}

func UserLogout(ctx context.Context) string {
	// 获取请求头中的token
	token := ctx.Value("token").(string)
	if err := services.RemoveToken(token); err != nil {
		return utils.ErrorResponse(err.Error())
	} else {
		return utils.SuccessResponse("登出成功", "success")
	}
}
