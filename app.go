/*
 * @Author: iRorikon
 * @Date: 2022-06-22 14:22:01
 * @FilePath: \ipsearch-go\app.go
 */
package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/nekomi-cn/ipsearch-go/config"
	"github.com/nekomi-cn/ipsearch-go/server"
)

func init() {
	flag.BoolVar(&config.Debug, "d", false, "start with debug mode")
	flag.BoolVar(&config.Debug, "debug", false, "start with debug mode")
	flag.BoolVar(&config.Version, "v", false, "print version info")
	flag.BoolVar(&config.Version, "version", false, "print version info")
	flag.Usage = usage
}

var usageStr = `
Usage: IPsearch [option]

Optionnal parameters:
    -v, --version                    Show software version
    -d, --debug                      Use debug mode
    -h, --help                       Show help message
`

func usage() {
	fmt.Println(usageStr)
	os.Exit(57)
}

func Init() bool {
	config.InitConfig()
	config.InitLog()
	config.InitDBFiles()
	return true
}

func main() {
	flag.Parse()
	if config.Version {
		fmt.Printf("APP Name: %s\nVersion: %s\n", config.AppName, config.VERSION)
		return
	}
	if !Init() {
		return
	}
	if config.Debug {
		log.Info("Set Debug Mode")
	}
	log.Info("Starting Client")
	config.PrintLogo()
	server.Start()
}
