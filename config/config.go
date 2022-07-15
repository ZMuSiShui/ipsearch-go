/*
 * @Author: iRorikon
 * @Date: 2022-06-22 14:22:01
 * @FilePath: \ipsearch-go\config\config.go
 */
package config

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/kirinlabs/HttpRequest"
	"github.com/nekomi-cn/ipsearch-go/util"
	log "github.com/sirupsen/logrus"
)

// 常量定义
const (
	AppName   string = "IPSearch-backend-go"
	VERSION   string = "1.2"
	UserAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	configURL string = "https://ipsearch.irorikon.workers.dev/"
)

// 变量定义

var (
	BuiltAt   string
	GoVersion string
)

var (
	CFG             *Config
	Debug           bool
	Version         bool
	UpdateDB        bool
	IPIPFile        string
	MaxmindFile     string
	CZ88File        string
	IP2LocationFile string
)

type Config struct {
	System SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
	DBList []DBConfig   `mapstructure:"db_list" json:"db_list" yaml:"db_list"`
}
type SystemConfig struct {
	ENV      string `mapstructure:"env" json:"env" yaml:"env"`
	DataPath string `mapstructure:"data_path" json:"data_path" yaml:"data_path"`
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
}

type DBConfig struct {
	DBFileName string `mapstructure:"file_name" json:"file_name" yaml:"file_name"`
	URL        string `mapstructure:"url" json:"url" yaml:"url"`
	MD5        string `mapstructure:"md5" json:"md5" yaml:"md5"`
}

func InitConfig() {
	req := HttpRequest.NewRequest()
	req.SetTimeout(30)
	req.SetHeaders(map[string]string{
		"User-Agent": UserAgent,
	})
	resp, err := req.Get(configURL, nil)
	if err != nil {
		panic(err)
	}
	defer resp.Close()
	err = resp.Json(&CFG)
	if err != nil {
		panic(err)
	}
}

func InitDBFiles() {
	for _, j := range CFG.DBList {
		log.Infof("Check DB File: %s", j.DBFileName)
		setFileName(fmt.Sprintf("%s/%s", CFG.System.DataPath, j.DBFileName))
		CheckIPDBFlie(fmt.Sprintf("%s/%s", CFG.System.DataPath, j.DBFileName), j.URL, j.MD5)
	}
}

func CheckIPDBFlie(path string, ipdburl string, md5 string) bool {
	if !util.FileExists(path) {
		log.Infof("IP DB file not exists, Download default IPDB file")
		err := util.StartDownload(ipdburl, path, runtime.NumCPU(), md5)
		if err != nil {
			log.Fatalf("Failed to download ipdb file. Error: %s", err)
		}
		return true
	}
	return true
}

func setFileName(i string) {
	if strings.Contains(i, "ipip") {
		IPIPFile = i
	} else if strings.Contains(i, "maxmind") {
		MaxmindFile = i
	} else if strings.Contains(i, "ip2location") {
		IP2LocationFile = i
	} else if strings.Contains(i, "cz88") {
		CZ88File = i
	}
}

func PrintLogo() {
	fmt.Print(`
                   _ooOoo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  O\  =  /O
               ____/'---'\____
             .'  \\|     |//  '.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||-  \
           |   | \\\  -  /// |   |
           | \_|  ''\-/''  |   |
           \  .-\__  '-'  ___/-. /
         ___'. .'  /-.-\  '. . __
      ."" '<  '.___\_<|>_/___.'  >'"".
     | | :  '- \'.;'\ _ /';.'/ - ' : | |
     \  \ '-.   \_ __\ /__ _/   .-' /  /
======'-.____'-.___\_____/___.-'____.-'======
                   '=-='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
            佛祖保佑       永无BUG
`)
}
