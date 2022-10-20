package project_detail

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (projectDetail ProjectDetail) {
    database.DB.Where("id", id).First(&projectDetail)
    return
}

func GetBy(field, value string) (projectDetail ProjectDetail) {
    database.DB.Where("? = ?", field, value).First(&projectDetail)
    return
}

func All() (projectDetails []ProjectDetail) {
    database.DB.Find(&projectDetails)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ProjectDetail{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (projectDetails []ProjectDetail, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ProjectDetail{}),
        &projectDetails,
        app.V1URL(database.TableName(&ProjectDetail{})),
        perPage,
    )
    return
}