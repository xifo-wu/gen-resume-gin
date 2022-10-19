// Package work_experience_detail 模型
package work_experience_detail

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
)

type WorkExperienceDetail struct {
	models.BaseModel

	Name             string `json:"name"`
	StartOn          string `json:"startOn"`
	EndOn            string `json:"endOnf"`
	Desc             string `json:"desc"`
	JobTitle         string `json:"jobTitle"`
	WorkExperienceID string `json:"workExperienceID"`
	SortIndex        int    `json:"sortIndex"`

	models.CommonTimestampsField
}

func (workExperienceDetail *WorkExperienceDetail) Create() {
	database.DB.Create(&workExperienceDetail)
}

func (workExperienceDetail *WorkExperienceDetail) Save() (rowsAffected int64) {
	result := database.DB.Save(&workExperienceDetail)
	return result.RowsAffected
}

func (workExperienceDetail *WorkExperienceDetail) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&workExperienceDetail)
	return result.RowsAffected
}
