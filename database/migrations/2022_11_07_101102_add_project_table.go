package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Project struct {
		models.BaseModel

		Label           string
		Visible         bool
		Config          string `gorm:"type:text"`
		ContentType     string
		ModuleTitleType string
		ResumeID        uint64 `gorm:"index"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Project{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Project{})
	}

	migrate.Add("2022_11_07_101102_add_project_table", up, down)
}
