package config

type Server struct {
	Port          int    `mapstructure:"port"`            // 端口
	Pprof         bool   `mapstructure:"pprof"`           // 是否开启 pprof 性能监控分析
	UseJsonNumber bool   `mapstructure:"use_json_number"` // 设置 true 以调用 JSON Decoder 实例上的 UseNumber 方法
	AppName       string `mapstructure:"app_name"`        // 应用名称 用于链路追踪注册标记
}
