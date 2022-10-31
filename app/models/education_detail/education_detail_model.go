// Package education_detail 模型
package education_detail

import (
	"backend/app/models"
	"backend/pkg/database"
)

type EducationDetail struct {
	models.BaseModel

	Name             string `json:"name"`
	StartOn          string `json:"startOn"`
	EndOn            string `json:"endOn"`
	UniversityMajors string `json:"universityMajors"`
	Desc             string `json:"desc"`
	EducationID      string `json:"educationID"`
	SortIndex        int    `json:"sortIndex"`

	models.CommonTimestampsField
}

func (educationDetail *EducationDetail) Create() {
	database.DB.Create(&educationDetail)
}

func (educationDetail *EducationDetail) Save() (rowsAffected int64) {
	result := database.DB.Save(&educationDetail)
	return result.RowsAffected
}

func (educationDetail *EducationDetail) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&educationDetail)
	return result.RowsAffected
}
