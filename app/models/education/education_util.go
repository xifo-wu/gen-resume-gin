package education

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (education Education) {
    database.DB.Where("id", id).First(&education)
    return
}

func GetBy(field, value string) (education Education) {
    database.DB.Where("? = ?", field, value).First(&education)
    return
}

func All() (educations []Education) {
    database.DB.Find(&educations)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Education{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (educations []Education, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Education{}),
        &educations,
        app.V1URL(database.TableName(&Education{})),
        perPage,
    )
    return
}