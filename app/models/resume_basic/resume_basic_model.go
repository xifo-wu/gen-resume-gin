// Package resume_basic 模型
package resume_basic

import (
	"backend/app/models"
	"backend/app/models/resume_basic_field_config"
	"backend/pkg/database"
)

type ResumeBasicFieldConfig = resume_basic_field_config.ResumeBasicFieldConfig

type ResumeBasic struct {
	models.BaseModel

	ResumeID                          string                 `json:"-"`
	Age                               int                    `json:"age"`                       // 年龄
	Birthday                          string                 `json:"birthday"`                  // 生日
	Avatar                            string                 `json:"avatar"`                    // 头像
	Email                             string                 `json:"email"`                     // 邮箱
	Job                               string                 `json:"job"`                       // 求职岗位
	Mobile                            string                 `json:"mobile"`                    // 手机号
	Name                              string                 `json:"name"`                      // 姓名
	Website                           string                 `json:"website"`                   // 个人网站
	EducationalQualifications         string                 `json:"educationalQualifications"` // 学历
	AgeConfig                         ResumeBasicFieldConfig `gorm:"foreignKey:AgeConfigID" json:"ageConfig"`
	BirthdayConfig                    ResumeBasicFieldConfig `gorm:"foreignKey:BirthdayConfigID" json:"birthdayConfig"`
	AvatarConfig                      ResumeBasicFieldConfig `gorm:"foreignKey:AvatarConfigID" json:"avatarConfig"`
	EmailConfig                       ResumeBasicFieldConfig `gorm:"foreignKey:EmailConfigID" json:"emailConfig"`
	JobConfig                         ResumeBasicFieldConfig `gorm:"foreignKey:JobConfigID" json:"jobConfig"`
	MobileConfig                      ResumeBasicFieldConfig `gorm:"foreignKey:MobileConfigID" json:"mobileConfig"`
	NameConfig                        ResumeBasicFieldConfig `gorm:"foreignKey:NameConfigID" json:"nameConfig"`
	WebsiteConfig                     ResumeBasicFieldConfig `gorm:"foreignKey:WebsiteConfigID" json:"websiteConfig"`
	EducationalQualificationsConfig   ResumeBasicFieldConfig `gorm:"foreignKey:EducationalQualificationsConfigID" json:"educationalQualificationsConfig"`
	AgeConfigID                       uint64
	BirthdayConfigID                  uint64
	AvatarConfigID                    uint64
	EmailConfigID                     uint64
	JobConfigID                       uint64
	MobileConfigID                    uint64
	NameConfigID                      uint64
	WebsiteConfigID                   uint64
	EducationalQualificationsConfigID uint64

	models.CommonTimestampsField
}

func (resumeBasic *ResumeBasic) Create() {
	database.DB.Create(&resumeBasic)
}

func (resumeBasic *ResumeBasic) Save() (rowsAffected int64) {
	result := database.DB.Save(&resumeBasic)
	return result.RowsAffected
}

func (resumeBasic *ResumeBasic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&resumeBasic)
	return result.RowsAffected
}
