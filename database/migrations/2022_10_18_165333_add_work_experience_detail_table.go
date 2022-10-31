package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type WorkExperienceDetail struct {
		models.BaseModel

		Name             string
		StartOn          string
		EndOn            string
		JobTitle         string
		Desc             string `gorm:"type:text"`
		WorkExperienceID string
		SortIndex        int `gorm:"default:0"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&WorkExperienceDetail{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&WorkExperienceDetail{})
	}

	migrate.Add("2022_10_18_165333_add_work_experience_detail_table", up, down)
}
