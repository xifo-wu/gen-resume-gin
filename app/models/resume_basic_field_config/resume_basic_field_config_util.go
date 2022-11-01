package resume_basic_field_config

import (
    "backend/pkg/app"
    "backend/pkg/database"
    "backend/pkg/paginator"

    "github.com/gin-gonic/gin"
)

func Get(id string) (resumeBasicFieldConfig ResumeBasicFieldConfig) {
    database.DB.Where("id", id).First(&resumeBasicFieldConfig)
    return
}

func GetBy(field, value string) (resumeBasicFieldConfig ResumeBasicFieldConfig) {
    database.DB.Where("? = ?", field, value).First(&resumeBasicFieldConfig)
    return
}

func All() (resumeBasicFieldConfigs []ResumeBasicFieldConfig) {
    database.DB.Find(&resumeBasicFieldConfigs)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(ResumeBasicFieldConfig{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (resumeBasicFieldConfigs []ResumeBasicFieldConfig, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(ResumeBasicFieldConfig{}),
        &resumeBasicFieldConfigs,
        app.V1URL(database.TableName(&ResumeBasicFieldConfig{})),
        perPage,
    )
    return
}