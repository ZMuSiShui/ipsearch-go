package conf

// 常量定义
const (
	AppName            string = "IPSearch-backend-go"
	VERSION            string = "1.1"
	IPIPFile           string = "data/ipip.ipdb"
	IPIPFileMD5        string = "ff82f6c47564bcf2374fb1f91c542795"
	IPIPURL            string = "https://eu.191013.xyz/src/ipip.ipdb"
	MaxmindFile        string = "data/maxmind.mmdb"
	MaxmindFileMD5     string = "cd7e60f57ae31d6b69387edf7b2064d7"
	MaxmindURL         string = "https://eu.191013.xyz/src/maxmind.mmdb"
	IP2LocationFile    string = "data/ip2location.bin"
	IP2LocationFileMD5 string = "64cfc677bab21e52f707ae2961044036"
	IP2LocationURL     string = "https://eu.191013.xyz/src/ip2location.bin"
	CZ88File           string = "data/cz88.dat"
	CZ88URL            string = "https://eu.191013.xyz/src/cz88.dat"
	CZ88FileMD5        string = "e551d51971a381ad4f7483aff0a23f9c"
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
