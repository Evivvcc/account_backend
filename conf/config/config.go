package config

type AppConfig struct {
	Server     Server                  `mapstructure:"server"` // 端口配置
	MultiDB    map[string]GeneralDB    `mapstructure:"db"`     // 数据库配置
	MultiRedis map[string]GeneralRedis `mapstructure:"redis"`  // Redis 配置
	Logger     Logger                  `mapstructure:"logger"` // 日志配置
	Http       Http                    `mapstructure:"http"`   // 三方配置
	RsaConfig  Rsa                     `mapstructure:"rsa"`    // RSA 密钥配置
}

// OneConfig 所有配置都挂载在该全局变量下，在 viper 文件的 init 中初始化该变量
var OneConfig AppConfig
