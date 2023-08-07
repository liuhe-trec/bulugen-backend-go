package conf

import (
	"bulugen-backend-go/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	// 设置数据库相关内容
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 同步表结构
	db.AutoMigrate(&model.User{})
	return db, nil
}
