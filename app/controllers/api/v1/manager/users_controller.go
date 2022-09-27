package manager

import (
	v1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/models/user"
	"gen-resume/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	v1.BaseAPIController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data": data,
		"meta": pager,
	})
}
