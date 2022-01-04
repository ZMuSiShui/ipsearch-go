package conf

// 常量定义
const (
	AppName         string = "IPSearch-backend-go"
	VERSION         string = "1.1"
	IPIPFile        string = "data/ipip.ipdb"
	IPIPURL         string = "https://eu.191013.xyz/src/ipip.ipdb"
	MaxmindFile     string = "data/maxmind.mmdb"
	MaxmindURL      string = "https://eu.191013.xyz/src/maxmind.mmdb"
	IP2LocationFile string = "data/ip2location.bin"
	IP2LocationURL  string = "https://eu.191013.xyz/src/ip2location.bin"
)

// 变量定义

var (
	BuiltAt   string
	GoVersion string
)

var (
	Debug      bool
	Version    bool
	SkipUpdate bool
)
