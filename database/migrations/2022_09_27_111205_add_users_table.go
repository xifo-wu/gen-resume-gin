package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Username string `gorm:"type:varchar(255);not null;index"`
		Nickname string `gorm:"type:varchar(255)"`
		Avatar   string `gorm:"type:text"`
		Email    string `gorm:"type:varchar(255)"`
		Phone    string `json:"type:varchar(20)"`
		Password string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_09_27_111205_add_users_table", up, down)
}
