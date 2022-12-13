package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasics struct {
		JobYearID uint64
		OrderKeys string
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasics{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&ResumeBasics{}, "JobYearID")
		migrator.DropColumn(&ResumeBasics{}, "OrderKeys")
	}

	migrate.Add("2022_12_13_151516_add_job_year_id_to_resume_basic", up, down)
}
