// Package models 模型通用属性和方法
package models

import (
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updatedAt,omitempty"`
}

type SoftDestroyField struct {
	DeletedAt time.Time `gorm:"column:deleted_at;index;" json:"deletedAt,omitempty"`
	IsDeleted bool      `gorm:"column:is_deleted;index;" json:"isDeleted,omitempty"`
}
