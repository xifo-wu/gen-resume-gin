package migrations

import (
	"database/sql"
	"gen-resume/app/models"
	"gen-resume/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Kv struct {
		models.BaseModel

		IsShowLabel bool
		Label       string
		Value       string
		Visible     bool
		Icon        string
		OwnerID     uint64
		OwnerType   string

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Kv{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Kv{})
	}

	migrate.Add("2022_10_05_180929_kv", up, down)
}
