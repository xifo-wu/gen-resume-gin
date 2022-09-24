package v1

import (
	"fmt"
	"gen-resume/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// PhoneIsExist
func (controller *UsersController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "无法解析" + err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
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
