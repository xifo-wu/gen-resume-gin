package auth

import (
	"errors"
	"gen-resume/app/models/user"
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
