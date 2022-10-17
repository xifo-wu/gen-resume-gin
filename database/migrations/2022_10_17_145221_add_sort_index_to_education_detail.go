package migrations

import (
	"database/sql"
	"gen-resume/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type EducationDetail struct {
		SortIndex int `gorm:"default:0"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&EducationDetail{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&EducationDetail{}, "ModuleOrder")
	}

	migrate.Add("2022_10_17_145221_add_sort_index_to_education_detail", up, down)
}
