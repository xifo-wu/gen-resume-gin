package requests

import (
	"backend/app/requests/validators"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thedevsaddam/govalidator"
)

// SignUpUsingPhoneRequest 通过手机注册的请求信息
type SignUpUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verifyCode,omitempty" valid:"verifyCode"`
	Username        string `valid:"username" json:"username"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"passwordConfirm" json:"passwordConfirm,omitempty"`
}

func SignUpUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":           []string{"required", "digits:11", "not_exists:users,phone"},
		"username":        []string{"required", "alpha_num", "between:3,32", "not_exists:users,username"},
		"password":        []string{"required", "min:6"},
		"passwordConfirm": []string{"required"},
		"verifyCode":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"username": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~32 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"passwordConfirm": []string{
			"required:确认密码为必填项",
		},
		"verifyCode": []string{
			"required:验证码必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignUpUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

// SignUpUsingEmailRequest 通过邮箱注册的请求信息
type SignUpUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verifyCode,omitempty" valid:"verifyCode"`
	Username        string `valid:"username" json:"username"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"passwordConfirm" json:"passwordConfirm,omitempty"`
}

func SignUpUsingEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":           []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"username":        []string{"required", "alpha_num", "between:3,32", "not_exists:users,username"},
		"password":        []string{"required", "min:6"},
		"passwordConfirm": []string{"required"},
		"verifyCode":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"username": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~32 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"passwordConfirm": []string{
			"required:确认密码框为必填项",
		},
		"verifyCode": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	_data := data.(*SignUpUsingEmailRequest)

	if len(_data.Username) == 0 {
		newID := uuid.New().String()

		_data.Username = strings.Replace(newID, "-", "", -1)
	}

	errs := validate(_data, rules, messages)

	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
