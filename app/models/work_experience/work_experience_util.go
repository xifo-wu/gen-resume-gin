package work_experience

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (workExperience WorkExperience) {
    database.DB.Where("id", id).First(&workExperience)
    return
}

func GetBy(field, value string) (workExperience WorkExperience) {
    database.DB.Where("? = ?", field, value).First(&workExperience)
    return
}

func All() (workExperiences []WorkExperience) {
    database.DB.Find(&workExperiences)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(WorkExperience{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (workExperiences []WorkExperience, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(WorkExperience{}),
        &workExperiences,
        app.V1URL(database.TableName(&WorkExperience{})),
        perPage,
    )
    return
}