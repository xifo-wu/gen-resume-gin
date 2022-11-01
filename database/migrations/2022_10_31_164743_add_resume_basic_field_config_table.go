package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasicFieldConfig struct {
		models.BaseModel

		ResumeBasicID uint64 `gorm:"index"`
		Visible       bool
		Label         string
		Icon          string
		ShowLabel     bool

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasicFieldConfig{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ResumeBasicFieldConfig{})
	}

	migrate.Add("2022_10_31_164743_add_resume_basic_field_config_table", up, down)
}
