// Package user 存放用户 Model 相关逻辑
package user

import (
	"gen-resume/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
