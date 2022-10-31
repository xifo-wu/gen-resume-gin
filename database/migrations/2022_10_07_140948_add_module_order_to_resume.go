package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Resume struct {
		ModuleOrder string
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Resume{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&Resume{}, "ModuleOrder")
	}

	migrate.Add("2022_10_07_140948_add_module_order_to_resume", up, down)
}
