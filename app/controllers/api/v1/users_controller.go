package v1

import (
	"backend/app/models/user"
	"backend/app/requests"
	"backend/pkg/auth"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// IsPhoneExist
func (controller *UsersController) IsPhoneExist(c *gin.Context) {
	request := requests.UsersPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.UsersPhoneExist); !ok {
		return
	}

	response.Data(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist
func (controller *UsersController) IsEmailExist(c *gin.Context) {
	request := requests.UsersEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.UsersEmailExist); !ok {
		return
	}

	response.Data(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Username, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}
