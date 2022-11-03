package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasic struct {
		AgeConfigID                       uint64
		BirthdayConfigID                  uint64
		AvatarConfigID                    uint64
		EmailConfigID                     uint64
		JobConfigID                       uint64
		MobileConfigID                    uint64
		NameConfigID                      uint64
		WebsiteConfigID                   uint64
		EducationalQualificationsConfigID uint64
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&ResumeBasic{}, "Gravatar")
	}

	migrate.Add("2022_11_02_161312_add_config_id_to_resume_basic", up, down)
}
