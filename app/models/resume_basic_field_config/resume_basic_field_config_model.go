// Package resume_basic_field_config 模型
package resume_basic_field_config

import (
	"backend/app/models"
	"backend/pkg/database"
)

type ResumeBasicFieldConfig struct {
	models.BaseModel

	ResumeBasicID uint64 `json:"-"`
	Visible       bool   `json:"visible"`
	Label         string `json:"label"`
	Icon          string `json:"icon"`
	ShowLabel     bool   `json:"showLabel"`
}

func (resumeBasicFieldConfig *ResumeBasicFieldConfig) Create() {
	database.DB.Create(&resumeBasicFieldConfig)
}

func (resumeBasicFieldConfig *ResumeBasicFieldConfig) Save() (rowsAffected int64) {
	result := database.DB.Save(&resumeBasicFieldConfig)
	return result.RowsAffected
}

func (resumeBasicFieldConfig *ResumeBasicFieldConfig) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&resumeBasicFieldConfig)
	return result.RowsAffected
}
