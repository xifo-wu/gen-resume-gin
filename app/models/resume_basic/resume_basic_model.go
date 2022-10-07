// Package resume_basic 模型
package resume_basic

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
)

type ResumeBasic struct {
	models.BaseModel

	ResumeID                        uint64 `json:"resumeID"`
	Age                             int    `json:"age"`                       // 年龄
	Birthday                        string `json:"birthday"`                  // 生日
	Avatar                          string `json:"avatar"`                    // 头像
	Email                           string `json:"email"`                     // 邮箱
	Job                             string `json:"job"`                       // 求职岗位
	Mobile                          string `json:"mobile"`                    // 手机号
	Name                            string `json:"name"`                      // 姓名
	Website                         string `json:"website"`                   // 个人网站
	EducationalQualifications       string `json:"educationalQualifications"` // 学历
	AgeConfig                       string `json:"ageConfig" gorm:"type:text"`
	BirthdayConfig                  string `json:"birthdayConfig" gorm:"type:text"`
	AvatarConfig                    string `json:"avatarConfig" gorm:"type:text"`
	EmailConfig                     string `json:"emailConfig" gorm:"type:text"`
	JobConfig                       string `json:"jobConfig" gorm:"type:text"`
	MobileConfig                    string `json:"mobileConfig" gorm:"type:text"`
	NameConfig                      string `json:"nameConfig" gorm:"type:text"`
	WebsiteConfig                   string `json:"websiteConfig" gorm:"type:text"`
	EducationalQualificationsConfig string `json:"educationalQualificationsConfig" gorm:"type:text"`

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
