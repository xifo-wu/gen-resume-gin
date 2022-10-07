// Package resume 模型
package resume

import (
	"gen-resume/app/models"
	"gen-resume/app/models/resume_basic"
	"gen-resume/app/models/user"
	"gen-resume/pkg/database"
)

type Resume struct {
	models.BaseModel

	Name        string                   `json:"name"`
	Slug        string                   `json:"slug"`
	LayoutType  string                   `json:"layoutType"`
	UserID      string                   `json:"-"`
	User        user.User                `json:"user"`
	ResumeBasic resume_basic.ResumeBasic `json:"resumeBasic"`

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
