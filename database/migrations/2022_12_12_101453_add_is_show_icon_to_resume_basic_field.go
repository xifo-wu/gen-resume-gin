package migrations

import (
	"backend/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type ResumeBasicField struct {
		ShowLabel   bool // 将要 Drop 的列
		IsShowIcon  bool `gorm:"default:false"`
		IsShowLabel bool `gorm:"default:true"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasicField{})
		migrator.DropColumn(&ResumeBasicField{}, "ShowLabel")
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AddColumn(&ResumeBasicField{}, "ShowLabel")
		migrator.DropColumn(&ResumeBasicField{}, "IsShowIcon")
	}

	migrate.Add("2022_12_12_101453_add_is_show_icon_to_resume_basic_field", up, down)
}
