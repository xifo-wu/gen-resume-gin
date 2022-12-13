// Package resume_basic_field 模型
package resume_basic_field

import (
	"backend/app/models"
	"backend/pkg/database"
)

type ResumeBasicField struct {
	models.BaseModel

	ResumeBasicID uint64 `json:"-"`
	Value         string `json:"value"`
	Visible       bool   `json:"visible"`
	Label         string `json:"label"`
	Icon          string `json:"icon"`
	IsShowLabel   bool   `json:"isShowLabel"`
	IsShowIcon    bool   `json:"isShowIcon"`
}

func (resumeBasicField *ResumeBasicField) Create() {
	database.DB.Create(&resumeBasicField)
}

func (resumeBasicField *ResumeBasicField) Save() (rowsAffected int64) {
	result := database.DB.Save(&resumeBasicField)
	return result.RowsAffected
}

func (resumeBasicField *ResumeBasicField) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&resumeBasicField)
	return result.RowsAffected
}
