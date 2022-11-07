package resume_basic_field

import (
	"backend/pkg/app"
	"backend/pkg/database"
	"backend/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(id string) (resumeBasicField ResumeBasicField) {
	database.DB.Where("id", id).First(&resumeBasicField)
	return
}

func GetBy(field, value string) (resumeBasicField ResumeBasicField) {
	database.DB.Where("? = ?", field, value).First(&resumeBasicField)
	return
}

func All() (resumeBasicFieldConfigs []ResumeBasicField) {
	database.DB.Find(&resumeBasicFieldConfigs)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(ResumeBasicField{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (resumeBasicFieldConfigs []ResumeBasicField, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(ResumeBasicField{}),
		&resumeBasicFieldConfigs,
		app.V1URL(database.TableName(&ResumeBasicField{})),
		perPage,
	)
	return
}
