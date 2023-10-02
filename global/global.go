package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// V viper 对象
var V *viper.Viper

// MysqlMap mysql.pool key is mysql database name
var MysqlMap = make(map[string]*gorm.DB)
