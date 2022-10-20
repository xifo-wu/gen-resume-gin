package migrations

import (
	"database/sql"
	"gen-resume/app/models"
	"gen-resume/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type ProjectDetail struct {
		models.BaseModel
		Name      string
		StartOn   string
		EndOn     string
		Role      string
		Desc      string
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
