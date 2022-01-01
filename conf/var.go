package conf

import "github.com/eko/gocache/v2/cache"

// 变量定义

var (
	BuiltAt   string
	GoVersion string
)

var (
	ConfigFile string
	Conf       *Config

	Debug      bool
	Version    bool
	SkipUpdate bool

	Cache *cache.Cache
)
