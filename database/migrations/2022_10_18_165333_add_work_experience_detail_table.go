package migrations

import (
	"database/sql"
	"gen-resume/app/models"
	"gen-resume/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type WorkExperienceDetail struct {
		models.BaseModel

		Name             string
		StartOn          string
		EndOn            string
		JobTitle         string
		Desc             string
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
