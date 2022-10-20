package project

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (project Project) {
    database.DB.Where("id", id).First(&project)
    return
}

func GetBy(field, value string) (project Project) {
    database.DB.Where("? = ?", field, value).First(&project)
    return
}

func All() (projects []Project) {
    database.DB.Find(&projects)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Project{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (projects []Project, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Project{}),
        &projects,
        app.V1URL(database.TableName(&Project{})),
        perPage,
    )
    return
}