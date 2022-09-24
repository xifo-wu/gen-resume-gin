// Package database 数据库操作
package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SqlDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 SqlDB
	SqlDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
