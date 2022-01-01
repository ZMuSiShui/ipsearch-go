package util

import (
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DownloadIPDB(filepath string, ipdburl string) (err error) {
	log.Infof("开始下载，目标地址: %s", ipdburl)

	res, err := http.Get(ipdburl)
	if err != nil {
		return
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = Write(filepath, b)
	if err != nil {
		return
	}
	return
}

func Write(path string, src []byte) (err error) {
	var file *os.File
	if FileExists(path) {
		file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return
		}
	} else {
		file, err = CreatNestedFile(path)
		if err != nil {
			return
		}
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write(src)
	if err != nil {
		return
	}
	return
}
