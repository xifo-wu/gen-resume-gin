package auth

import (
	v1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/models/user"
	"gen-resume/app/requests"
	"gen-resume/pkg/jwt"
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
	userModel := user.User{
		Username: request.Username,
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Username)
		response.CreatedJSON(c, gin.H{
			"meta": gin.H{
				"token": token,
			},
			"data": userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}

// SignUpUsingEmail 使用 Email + 验证码进行注册
func (sc *SignUpController) SignUpUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignUpUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignUpUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Username)
		response.CreatedJSON(c, gin.H{
			"meta": gin.H{
				"token": token,
			},
			"data": userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
