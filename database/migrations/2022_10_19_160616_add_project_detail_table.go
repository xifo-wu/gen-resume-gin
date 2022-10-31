package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ProjectDetail struct {
		models.BaseModel
		Name      string
		StartOn   string
		EndOn     string
		Role      string
		Desc      string `gorm:"type:text"`
		ProjectID string
		SortIndex int `gorm:"default:0"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ProjectDetail{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ProjectDetail{})
	}

	migrate.Add("2022_10_19_160616_add_project_detail_table", up, down)
}
