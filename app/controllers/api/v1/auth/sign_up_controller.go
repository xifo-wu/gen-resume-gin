package auth

import (
	v1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/models/user"
	"gen-resume/app/requests"
	"gen-resume/pkg/response"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	v1.BaseAPIController
}

// SignUpUsingPhone 使用手机和验证码进行注册
func (sc *SignUpController) SignUpUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignUpUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignUpUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	_user := user.User{
		Username: request.Username,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
