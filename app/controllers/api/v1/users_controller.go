package v1

import (
	"fmt"
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

	// 表单验证
	errs := requests.ValidateUsersPhoneExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 只返回第一个错误
		for _, v := range errs {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": v[0],
			})
			return
		}

		// // 验证失败，返回 422 状态码和错误信息
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"errors": errs,
		// })
		// return
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

	// 表单验证
	errs := requests.ValidateUsersEmailExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 只返回第一个错误
		for _, v := range errs {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": v[0],
			})
			return
		}

		// // 验证失败，返回 422 状态码和错误信息
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		// 	"errors": errs,
		// })
		// return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"exist": user.IsEmailExist(request.Email),
		},
	})
}
