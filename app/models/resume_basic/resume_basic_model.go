// Package resume_basic 模型
package resume_basic

import (
	"backend/app/models"
	"backend/app/models/resume_basic_field"
	"backend/pkg/database"
)

type ResumeBasicField = resume_basic_field.ResumeBasicField

type ResumeBasic struct {
	models.BaseModel

	ResumeID                    string           `json:"-"`
	AgeID                       uint64           `json:"-"`
	BirthdayID                  uint64           `json:"-"`
	AvatarID                    uint64           `json:"-"`
	EmailID                     uint64           `json:"-"`
	JobID                       uint64           `json:"-"`
	MobileID                    uint64           `json:"-"`
	NameID                      uint64           `json:"-"`
	WebsiteID                   uint64           `json:"-"`
	EducationalQualificationsID uint64           `json:"-"`
	Age                         ResumeBasicField `gorm:"foreignKey:AgeID" json:"age"`
	Birthday                    ResumeBasicField `gorm:"foreignKey:BirthdayID" json:"birthday"`
	Avatar                      ResumeBasicField `gorm:"foreignKey:AvatarID" json:"avatar"`
	Email                       ResumeBasicField `gorm:"foreignKey:EmailID" json:"email"`
	Job                         ResumeBasicField `gorm:"foreignKey:JobID" json:"job"`
	Mobile                      ResumeBasicField `gorm:"foreignKey:MobileID" json:"mobile"`
	Name                        ResumeBasicField `gorm:"foreignKey:NameID" json:"name"`
	Website                     ResumeBasicField `gorm:"foreignKey:WebsiteID" json:"website"`
	EducationalQualifications   ResumeBasicField `gorm:"foreignKey:EducationalQualificationsID" json:"educationalQualifications"`

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
