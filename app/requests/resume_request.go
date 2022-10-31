package requests

import (
	"backend/pkg/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResumeRequest struct {
	Name       string `valid:"name" json:"name"`
	Slug       string `valid:"slug" json:"slug"`
	LayoutType string `valid:"layoutType" json:"layoutType"`
}

func ResumeSave(data interface{}, c *gin.Context) map[string][]string {
	id := c.Param("id")

	fmt.Println("id", id)
	rules := govalidator.MapData{
		"name":       []string{"required"},
		"slug":       []string{"required", "not_exists_in_user:resumes,slug," + auth.CurrentUserID(c) + "," + id},
		"layoutType": []string{"required"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
		},
		"slug": []string{
			"required:标识为必填项",
			"not_exists_in_user:标识已被占用",
		},
		"layoutType": []string{
			"required:简历模版为必填项",
		},
	}

	return validate(data, rules, messages)
}

type ResumeBasicRequest struct {
	Name string `json:"name"`
}
