package migrations

import (
	"backend/app/models"
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasics struct {
		models.BaseModel

		ResumeID                    string `gorm:"index"`
		AgeID                       uint64
		BirthdayID                  uint64
		AvatarID                    uint64
		EmailID                     uint64
		JobID                       uint64
		MobileID                    uint64
		NameID                      uint64
		WebsiteID                   uint64
		EducationalQualificationsID uint64

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasics{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ResumeBasics{})
	}

	migrate.Add("2022_11_07_100335_add_resume_basic_table", up, down)
}
