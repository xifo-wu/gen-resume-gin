// Package project_detail 模型
package project_detail

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
)

type ProjectDetail struct {
	models.BaseModel

	Name      string `json:"name"`
	StartOn   string `json:"startOn"`
	EndOn     string `json:"endOn"`
	Desc      string `json:"desc"`
	Role      string `json:"role"`
	ProjectID string `json:"projectID"`
	SortIndex int    `json:"sortIndex"`

	models.CommonTimestampsField
}

func (projectDetail *ProjectDetail) Create() {
	database.DB.Create(&projectDetail)
}

func (projectDetail *ProjectDetail) Save() (rowsAffected int64) {
	result := database.DB.Save(&projectDetail)
	return result.RowsAffected
}

func (projectDetail *ProjectDetail) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&projectDetail)
	return result.RowsAffected
}
