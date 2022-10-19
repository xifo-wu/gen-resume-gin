package work_experience_detail

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (workExperienceDetail WorkExperienceDetail) {
    database.DB.Where("id", id).First(&workExperienceDetail)
    return
}

func GetBy(field, value string) (workExperienceDetail WorkExperienceDetail) {
    database.DB.Where("? = ?", field, value).First(&workExperienceDetail)
    return
}

func All() (workExperienceDetails []WorkExperienceDetail) {
    database.DB.Find(&workExperienceDetails)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(WorkExperienceDetail{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (workExperienceDetails []WorkExperienceDetail, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(WorkExperienceDetail{}),
        &workExperienceDetails,
        app.V1URL(database.TableName(&WorkExperienceDetail{})),
        perPage,
    )
    return
}