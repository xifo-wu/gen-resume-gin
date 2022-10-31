package auth

import (
	"backend/app/models/user"
	"backend/pkg/logger"
	"errors"

	"github.com/gin-gonic/gin"
)

func Attempt(login string, password string) (user.User, error) {
	userModel := user.GetByMulti(login)

	if userModel.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码不正确")
	}

	return userModel, nil
}

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}

func LoginByEmail(email string) (user.User, error) {
	userModel := user.GetByEmail(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("邮箱未注册")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUserID 从 gin.context 中获取当前登录用户 ID
func CurrentUserID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
