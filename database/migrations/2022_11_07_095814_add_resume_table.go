package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Resume struct {
		models.BaseModel

		Name         string `gorm:"type:varchar(255);not null;index"`
		Slug         string `gorm:"type:varchar(255);not null;index"`
		LayoutType   string `gorm:"type:varchar(255)"`
		ThemeColor   string
		CustomStyles string `gorm:"type:text"`
		UserID       uint64 `gorm:"index"`
		ModuleOrder  string

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Resume{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Resume{})
	}

	migrate.Add("2022_11_07_095814_add_resume_table", up, down)
}
