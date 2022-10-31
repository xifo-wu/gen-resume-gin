package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		Gravatar string `gorm:"type:varchar(255)"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "Gravatar")
	}

	migrate.Add("2022_09_29_100211_add_gravatar_to_user", up, down)
}
