package config

import "fmt"

type GeneralDB struct {
	Host              string `mapstructure:"host"`
	Dbname            string `mapstructure:"dbname"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	Config            string `mapstructure:"cinfig"`
	MaxLifetime       int    `mapstructure:"max_lifetime" default:"7200"`       // 设置连接可以重用的最长时间（单位：秒）
	MaxOpenConnection int    `mapstructure:"max_open_connection" default:"500"` // 打开到数据库的最大连接数
	MaxIdleConnection int    `mapstructure:"max_idle_connection" default:"100"` // 空闲中的最大连接数
}

func (g *GeneralDB) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s&loc=local",
		g.Username,
		g.Password,
		g.Host,
		g.Dbname,
		g.Config,
	)
}
