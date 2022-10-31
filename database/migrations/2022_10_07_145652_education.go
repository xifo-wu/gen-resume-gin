package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Education struct {
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
		migrator.AutoMigrate(&Education{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Education{})
	}

	migrate.Add("2022_10_07_145652_education", up, down)
}
