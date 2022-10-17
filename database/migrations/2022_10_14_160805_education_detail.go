package migrations

import (
	"database/sql"
	"gen-resume/app/models"
	"gen-resume/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type EducationDetail struct {
		models.BaseModel

		Name             string
		StartOn          string
		EndOn            string
		UniversityMajors string
		Desc             string
		EducationID      string

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&EducationDetail{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&EducationDetail{})
	}

	migrate.Add("2022_10_14_160805_education_detail", up, down)
}
