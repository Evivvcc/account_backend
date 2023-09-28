package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败： %s \n", err))
	}
}
