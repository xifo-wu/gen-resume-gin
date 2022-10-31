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

		ResumeID                        string `json:"resumeID"`
		Age                             int    // 年龄
		AgeConfig                       string `gorm:"type:text"`
		Birthday                        string // 生日
		BirthdayConfig                  string `gorm:"type:text"`
		Avatar                          string // 头像
		AvatarConfig                    string `gorm:"type:text"`
		Email                           string // 邮箱
		EmailConfig                     string `gorm:"type:text"`
		Job                             string // 求职岗位
		JobConfig                       string `gorm:"type:text"`
		Mobile                          string // 手机号
		MobileConfig                    string `gorm:"type:text"`
		Name                            string // 姓名
		NameConfig                      string `gorm:"type:text"`
		Website                         string // 个人网站
		WebsiteConfig                   string `gorm:"type:text"`
		EducationalQualifications       string // 学历
		EducationalQualificationsConfig string `gorm:"type:text"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ResumeBasics{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ResumeBasics{})
	}

	migrate.Add("2022_10_06_175242_resume_basic", up, down)
}
