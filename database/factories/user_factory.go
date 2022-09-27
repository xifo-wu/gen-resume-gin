// Package factories 存放工厂方法
package factories

import (
	"gen-resume/app/models/user"
	"gen-resume/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var users []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Username: faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		users = append(users, model)
	}

	return users
}
