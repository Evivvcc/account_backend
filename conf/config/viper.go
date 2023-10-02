package config

import (
	"account_backend/global"
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置文件
// 1. 读取配置文件到 map 中；2. map 转化为 struct
func InitConfig(filePath string) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败： %s \n", err))
	}
	err = viper.Unmarshal(&OneConfig)
	if err != nil {
		panic("配置获取失败")
	}
	global.V = viper.GetViper()
}
