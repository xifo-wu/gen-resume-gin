// Package project 模型
package project

import (
	"gen-resume/app/models"
	"gen-resume/app/models/project_detail"
	"gen-resume/pkg/database"
)

type Project struct {
	models.BaseModel

	Label           string                          `json:"label"`
	Visible         bool                            `json:"visible"`
	Config          string                          `json:"config"`
	ContentType     string                          `json:"contentType"`
	ModuleTitleType string                          `json:"moduleTitleType"`
	ResumeID        string                          `json:"resumeID"`
	ProjectDetails  []*project_detail.ProjectDetail `json:"projectDetails"`

	models.CommonTimestampsField
}

func (project *Project) Create() {
	database.DB.Create(&project)
}

func (project *Project) Save() (rowsAffected int64) {
	result := database.DB.Save(&project)
	return result.RowsAffected
}

func (project *Project) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&project)
	return result.RowsAffected
}
