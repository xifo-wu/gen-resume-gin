package other

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (other Other) {
    database.DB.Where("id", id).First(&other)
    return
}

func GetBy(field, value string) (other Other) {
    database.DB.Where("? = ?", field, value).First(&other)
    return
}

func All() (others []Other) {
    database.DB.Find(&others)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Other{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (others []Other, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Other{}),
        &others,
        app.V1URL(database.TableName(&Other{})),
        perPage,
    )
    return
}