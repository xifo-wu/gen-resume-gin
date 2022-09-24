package v1

import (
	"gen-resume/app/models/user"
	"gen-resume/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// IsPhoneExist
func (controller *UsersController) IsPhoneExist(c *gin.Context) {
	request := requests.UsersPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.UsersPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"exist": user.IsPhoneExist(request.Phone),
		},
	})
}

// IsEmailExist
func (controller *UsersController) IsEmailExist(c *gin.Context) {
	request := requests.UsersEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.UsersEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"exist": user.IsEmailExist(request.Email),
		},
	})
}
