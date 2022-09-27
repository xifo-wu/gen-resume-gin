package v1

import (
	"gen-resume/app/models/user"
	"gen-resume/app/requests"
	"gen-resume/pkg/auth"
	"gen-resume/pkg/response"

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
