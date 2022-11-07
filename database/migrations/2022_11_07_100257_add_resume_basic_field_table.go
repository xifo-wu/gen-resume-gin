package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasicField struct {
		models.BaseModel

		ResumeBasicID uint64 `gorm:"index"`
		Value         string
		Visible       bool
		Label         string
		Icon          string
		ShowLabel     bool

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasicField{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ResumeBasicField{})
	}

	migrate.Add("2022_11_07_100257_add_resume_basic_field_table", up, down)
}
