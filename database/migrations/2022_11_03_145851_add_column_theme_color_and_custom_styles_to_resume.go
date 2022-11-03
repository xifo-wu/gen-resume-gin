package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {
	type Resume struct {
		ThemeColor   string
		CustomStyles string `gorm:"type:text"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Resume{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&Resume{}, "")
	}

	migrate.Add("2022_11_03_145851_add_column_theme_color_and_custom_styles_to_resume", up, down)
}
