package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type EducationDetail struct {
		models.BaseModel

		Name             string
		StartOn          string
		EndOn            string
		UniversityMajors string
		Desc             string `gorm:"type:text"`
		EducationID      string
		SortIndex        int `gorm:"default:0"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&EducationDetail{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&EducationDetail{})
	}

	migrate.Add("2022_11_07_100654_add_education_detail_table", up, down)
}
