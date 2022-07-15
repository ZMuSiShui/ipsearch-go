/*
 * @Author: iRorikon
 * @Date: 2022-06-22 14:22:01
 * @FilePath: \ipsearch-go\config\log.go
 */
package config

import (
	log "github.com/sirupsen/logrus"
)

// 初始化日志
func InitLog() {
	if Debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
	log.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
	})
}
