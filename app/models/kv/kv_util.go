package kv

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (kv Kv) {
    database.DB.Where("id", id).First(&kv)
    return
}

func GetBy(field, value string) (kv Kv) {
    database.DB.Where("? = ?", field, value).First(&kv)
    return
}

func All() (kvs []Kv) {
    database.DB.Find(&kvs)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Kv{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (kvs []Kv, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Kv{}),
        &kvs,
        app.V1URL(database.TableName(&Kv{})),
        perPage,
    )
    return
}