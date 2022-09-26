// Package user 存放用户 Model 相关逻辑
package user

import (
	"gen-resume/app/models"
	"gen-resume/pkg/database"
	"gen-resume/pkg/hash"
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

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
