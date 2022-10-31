package manager

import (
	v1 "backend/app/controllers/api/v1"
	"backend/app/models/user"
	"backend/app/requests"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	v1.BaseAPIController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data": data,
		"meta": pager,
	})
}
