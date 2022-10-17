// Package education 模型
package education

import (
	"gen-resume/app/models"
	"gen-resume/app/models/education_detail"
	"gen-resume/pkg/database"
)

type Education struct {
	models.BaseModel

	Label            string                              `json:"label"`
	Visible          bool                                `json:"visible"`
	Config           string                              `json:"config"`
	ContentType      string                              `json:"contentType"`
	ModuleTitleType  string                              `json:"moduleTitleType"`
	ResumeID         string                              `json:"resumeID"`
	EducationDetails []*education_detail.EducationDetail `json:"educationDetails"`

	models.CommonTimestampsField
}

func (education *Education) Create() {
	database.DB.Create(&education)
}

func (education *Education) Save() (rowsAffected int64) {
	result := database.DB.Save(&education)
	return result.RowsAffected
}

func (education *Education) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&education)
	return result.RowsAffected
}
