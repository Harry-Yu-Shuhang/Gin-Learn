package dao

import (
	"gin-learn/config"
	"gin-learn/pkg/logger"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	Db  *gorm.DB //想在别的程序调用，只能全局声明,且大写开头
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{}) // 连接数据库
	//以下是mysql提供的高级配置方法
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	// 	DefaultStringSize: 256, // string 类型字段的默认长度
	// 	DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	// 	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	// 	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	// 	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//   }), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]any{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		logger.Error(map[string]any{"database error": Db.Error.Error()})
	}

	// 获取底层的 *sql.DB(gorm V2后DB()对象不可以直接用，要错误处理)
	sqlDB, err := Db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}

	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接最大存活时间,一小时

}
