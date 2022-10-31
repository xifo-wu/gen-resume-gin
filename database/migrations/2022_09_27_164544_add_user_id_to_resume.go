package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Resume struct {
		UserID int `gorm:"index"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AddColumn(&Resume{}, "UserID")
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&Resume{}, "UserID")
	}

	migrate.Add("2022_09_27_164544_add_user_id_to_resume", up, down)
}
