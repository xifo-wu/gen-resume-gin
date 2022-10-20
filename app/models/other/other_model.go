// Package other 模型
package other

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
)

type Other struct {
	models.BaseModel

	Label           string `json:"label"`
	Visible         bool   `json:"visible"`
	Config          string `json:"config"`
	ContentType     string `json:"contentType"`
	ModuleTitleType string `json:"moduleTitleType"`
	ResumeID        string `json:"resumeID"`
	Desc            string `json:"desc"`
	SortIndex       int    `json:"sortIndex"`

	models.CommonTimestampsField
}

func (other *Other) Create() {
	database.DB.Create(&other)
}

func (other *Other) Save() (rowsAffected int64) {
	result := database.DB.Save(&other)
	return result.RowsAffected
}

func (other *Other) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&other)
	return result.RowsAffected
}
