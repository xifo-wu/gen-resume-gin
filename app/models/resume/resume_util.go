package resume

import (
	"backend/pkg/app"
	"backend/pkg/database"
	"backend/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(id string) (resume Resume) {
	database.DB.Model(&Resume{}).Preload(clause.Associations).Where("id", id).First(&resume)
	return
}

func GetBy(field, value string) (resume Resume) {
	database.DB.Where("? = ?", field, value).First(&resume)
	return
}

func All() (resumes []Resume) {
	database.DB.Find(&resumes)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Resume{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (resumes []Resume, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Resume{}),
		&resumes,
		app.V1URL(database.TableName(&Resume{})),
		perPage,
	)
	return
}
