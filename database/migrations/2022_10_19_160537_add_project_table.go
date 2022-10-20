package migrations

import (
	"database/sql"
	"gen-resume/app/models"
	"gen-resume/pkg/migrate"

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

	migrate.Add("2022_10_19_160537_add_project_table", up, down)
}
