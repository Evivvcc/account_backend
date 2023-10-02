package mysql

import (
	"account_backend/conf/config"
	"account_backend/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var mysqlOnce sync.Once

// InitMysql 数据库初始化连接
func InitMysql() {
	if len(config.OneConfig.MultiDB) == 0 {
		panic("数据库配置未加载")
	}
	mysqlOnce.Do(func() {
		databaseConn()
	})
}

// databaseConn 建立数据库连接
func databaseConn() {
	// 根据配置文件依次初始化
	for k, v := range config.OneConfig.MultiDB {
		mysqlConfig := mysql.Config{
			DSN:                       v.Dsn(), // DSN data source name
			DefaultStringSize:         191,     // string 类型字段的默认长度
			DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,   // 根据版本自动配置
		}
		gormConfig := &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   false, // 跳过默认事务
			Logger:                                   logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单表
			},
		}
		if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
			panic(fmt.Errorf("数据库初始化异常：%s\n", err))
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxOpenConns(v.MaxOpenConnection)
			sqlDB.SetMaxIdleConns(v.MaxIdleConnection)
			sqlDB.SetConnMaxLifetime(time.Duration(v.MaxLifetime) * time.Second)

			global.MysqlMap[k] = db
		}
	}
}
