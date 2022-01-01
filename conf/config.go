package conf

// 配置文件定义

type Config struct {
	AppName string `json:"appname"`
}

func DefaultConfig() *Config {
	return &Config{
		AppName: "My-Go-Template",
	}
}
