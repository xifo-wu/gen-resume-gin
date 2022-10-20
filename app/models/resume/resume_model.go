// Package resume 模型
package resume

import (
	"gen-resume/app/models"
	"gen-resume/app/models/education"
	"gen-resume/app/models/project"
	"gen-resume/app/models/resume_basic"
	"gen-resume/app/models/user"
	"gen-resume/app/models/work_experience"
	"gen-resume/pkg/database"
)

type Resume struct {
	models.BaseModel

	Name           string                          `json:"name"`
	Slug           string                          `json:"slug"`
	LayoutType     string                          `json:"layoutType"`
	UserID         string                          `json:"-"`
	User           user.User                       `json:"user"`
	ModuleOrder    string                          `json:"moduleOrder"`
	ResumeBasic    *resume_basic.ResumeBasic       `json:"resumeBasic,omitempty"`
	Education      *education.Education            `json:"education,omitempty"`
	Project        *project.Project                `json:"project,omitempty"`
	WorkExperience *work_experience.WorkExperience `json:"workExperience,omitempty"`

	models.CommonTimestampsField
}

func (resume *Resume) Create() {
	database.DB.Create(&resume)
}

func (resume *Resume) Save() (rowsAffected int64) {
	result := database.DB.Save(&resume)
	return result.RowsAffected
}

func (resume *Resume) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&resume)
	return result.RowsAffected
}
