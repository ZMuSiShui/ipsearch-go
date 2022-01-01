package base

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ZMuSiShui/ipsearch-go/conf"
	"github.com/ZMuSiShui/ipsearch-go/util"
	log "github.com/sirupsen/logrus"
)

// 初始化配置文件
func InitConfig() {
	log.Infof("Reading config file: %s", conf.ConfigFile)
	if !util.FileExists(conf.ConfigFile) {
		log.Infof("Config file not exists, Creating default config file")
		_, err := util.CreatNestedFile(conf.ConfigFile)
		if err != nil {
			log.Fatalf("Failed to create config file")
		}
		conf.Conf = conf.DefaultConfig()
		if !util.WriteToJson(conf.ConfigFile, conf.Conf) {
			log.Fatalf("Failed to create default config file")
		}
		return
	}
	config, err := ioutil.ReadFile(conf.ConfigFile)
	if err != nil {
		log.Fatalf("Reading config file error:%s", err.Error())
	}
	conf.Conf = new(conf.Config)
	err = json.Unmarshal(config, conf.Conf)
	if err != nil {
		log.Fatalf("Load config error: %s", err.Error())
	}
	log.Debug("config:%+v", conf.Conf)
}
