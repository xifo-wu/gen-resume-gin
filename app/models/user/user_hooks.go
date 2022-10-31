package user

import (
	"backend/pkg/hash"
	"crypto/md5"
	"fmt"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	if len(userModel.Email) != 0 {
		dataMD5Sum := md5.Sum([]byte(userModel.Email))
		md5str := fmt.Sprintf("%x", dataMD5Sum[:])
		userModel.Gravatar = "https://www.gravatar.com/avatar/" + md5str
	}

	return
}
