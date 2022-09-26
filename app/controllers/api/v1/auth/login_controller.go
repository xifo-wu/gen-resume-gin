package auth

import (
	v1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/requests"
	"gen-resume/pkg/auth"
	"gen-resume/pkg/jwt"
	"gen-resume/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "用户不存在")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)

		response.JSON(c, gin.H{
			"data": user,
			"meta": gin.H{
				"token": token,
			},
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败，显示错误提示
		// response.Unauthorized(c, "账号不存在或密码错误")
		response.Unauthorized(c, err.Error())

	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Username)
		response.JSON(c, gin.H{
			"data": user,
			"meta": gin.H{
				"token": token,
			},
		})
	}
}