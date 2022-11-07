package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Other struct {
		models.BaseModel

		Label           string
		Visible         bool
		Config          string `gorm:"type:text"`
		ContentType     string
		ModuleTitleType string
		ResumeID        uint64 `gorm:"index"`
		Desc            string `gorm:"type:text"`
		SortIndex       int    `gorm:"default:0"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Other{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Other{})
	}

	migrate.Add("2022_11_07_101217_add_other_table", up, down)
}
