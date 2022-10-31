// Package work_experience 模型
package work_experience

import (
	"backend/app/models"
	"backend/app/models/work_experience_detail"
	"backend/pkg/database"
)

type WorkExperience struct {
	models.BaseModel

	Label                 string                                         `json:"label"`
	Visible               bool                                           `json:"visible"`
	Config                string                                         `json:"config"`
	ContentType           string                                         `json:"contentType"`
	ModuleTitleType       string                                         `json:"moduleTitleType"`
	ResumeID              string                                         `json:"resumeID"`
	WorkExperienceDetails []*work_experience_detail.WorkExperienceDetail `json:"workExperienceDetails"`

	models.CommonTimestampsField
}

func (workExperience *WorkExperience) Create() {
	database.DB.Create(&workExperience)
}

func (workExperience *WorkExperience) Save() (rowsAffected int64) {
	result := database.DB.Save(&workExperience)
	return result.RowsAffected
}

func (workExperience *WorkExperience) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&workExperience)
	return result.RowsAffected
}
