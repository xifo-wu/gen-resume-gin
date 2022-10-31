package education_detail

import (
	"backend/pkg/app"
	"backend/pkg/database"
	"backend/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(id string) (educationDetail EducationDetail) {
	database.DB.Where("id", id).First(&educationDetail)
	return
}

func GetBy(field, value string) (educationDetail EducationDetail) {
	database.DB.Where("? = ?", field, value).First(&educationDetail)
	return
}

func All() (educationDetails []EducationDetail) {
	database.DB.Find(&educationDetails)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(EducationDetail{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (educationDetails []EducationDetail, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(EducationDetail{}),
		&educationDetails,
		app.V1URL(database.TableName(&EducationDetail{})),
		perPage,
	)
	return
}
