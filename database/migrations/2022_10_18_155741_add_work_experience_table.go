package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type WorkExperience struct {
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
		migrator.AutoMigrate(&WorkExperience{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&WorkExperience{})
	}

	migrate.Add("2022_10_18_155741_add_work_experience_table", up, down)
}
