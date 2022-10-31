// Package requests 处理请求数据和表单验证
package requests

import (
	"backend/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UsersPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func UsersPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, messages)
}

type UsersEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func UsersEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}

type UserUpdatePasswordRequest struct {
	Password           string `valid:"password" json:"password,omitempty"`
	NewPassword        string `valid:"newPassword" json:"newPassword,omitempty"`
	NewPasswordConfirm string `valid:"newPasswordConfirm" json:"newPasswordConfirm,omitempty"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":           []string{"required", "min:6"},
		"newPassword":        []string{"required", "min:6"},
		"newPasswordConfirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"newPassword": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"newPasswordConfirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}

	// 确保 confirm 密码正确
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}
