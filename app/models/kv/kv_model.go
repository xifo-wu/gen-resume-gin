// Package kv 模型
package kv

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
)

type Kv struct {
	models.BaseModel

	IsShowLabel bool   `json:"isShowLabel"`
	Label       string `json:"label"`
	Value       string `json:"value"`
	Visible     bool   `json:"visible"`
	Icon        string `json:"icon"`
	OwnerID     uint64 `json:"ownerID"`
	OwnerType   string `json:"ownerType"`

	models.CommonTimestampsField
}

func (kv *Kv) Create() {
	database.DB.Create(&kv)
}

func (kv *Kv) Save() (rowsAffected int64) {
	result := database.DB.Save(&kv)
	return result.RowsAffected
}

func (kv *Kv) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&kv)
	return result.RowsAffected
}
