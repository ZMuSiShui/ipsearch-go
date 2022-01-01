package conf

// 常量定义
const (
	AppName     string = "IPSearch-backend-go"
	VERSION     string = "1.0"
	IPIPFile    string = "data/ipip.ipdb"
	IPIPURL     string = "https://eu.191013.xyz/src/ipip.ipdb"
	MaxmindFile string = "data/maxmind.mmdb"
	MaxmindURL  string = "https://eu.191013.xyz/src/maxmind.mmdb"
	CZ88File    string = "data/cz88.dat"
	CZ88URL     string = "https://eu.191013.xyz/src/cz88.dat"
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
