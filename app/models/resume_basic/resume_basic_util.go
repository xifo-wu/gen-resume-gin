package resume_basic

import (
    "gen-resume/pkg/app"
    "gen-resume/pkg/database"
    "gen-resume/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (resumeBasic ResumeBasic) {
    database.DB.Where("id", id).First(&resumeBasic)
    return
}

func GetBy(field, value string) (resumeBasic ResumeBasic) {
    database.DB.Where("? = ?", field, value).First(&resumeBasic)
    return
}

func All() (resumeBasics []ResumeBasic) {
    database.DB.Find(&resumeBasics)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ResumeBasic{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (resumeBasics []ResumeBasic, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ResumeBasic{}),
        &resumeBasics,
        app.V1URL(database.TableName(&ResumeBasic{})),
        perPage,
    )
    return
}