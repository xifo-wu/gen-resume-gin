package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasic struct {
		InAWord string `gorm:"type:varchar(255)"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&ResumeBasic{}, "InAWord")
	}

	migrate.Add("2022_12_12_101424_add_in_a_word_to_resume_basic", up, down)
}
