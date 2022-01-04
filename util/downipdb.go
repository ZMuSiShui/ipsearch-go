package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

const userAgent = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36`

type fileDownload struct {
	fileSize     int
	url          string
	outFileName  string
	totalPart    int
	doneFilePart []filePart
	md5          string
}

type filePart struct {
	Index int
	From  int
	To    int
	Data  []byte
}

// 构建下载工厂函数
func NewFileDownload(url, outFileName string, totalPart int, md5 string) (*fileDownload, error) {
	return &fileDownload{
		fileSize:     0,
		url:          url,
		outFileName:  outFileName,
		totalPart:    totalPart,
		doneFilePart: make([]filePart, totalPart),
		md5:          md5,
	}, nil
}

// 开始下载任务
func StartDownload(url, outFileName string, totalPart int, md5 string) (err error) {
	downloader, err := NewFileDownload(url, outFileName, totalPart, md5)
	if err != nil {
		return
	}
	if err := downloader.Download(); err != nil {
		return err
	}
	return nil
}

// 下载主程序
func (d *fileDownload) Download() (err error) {
	fileTotalSize, err := d.getHeaderInfo()
	if err != nil {
		return
	}
	d.fileSize = fileTotalSize

	jobs := make([]filePart, d.totalPart)
	eachSize := fileTotalSize / d.totalPart

	for i := range jobs {
		jobs[i].Index = i
		if i == 0 {
			jobs[i].From = 0
		} else {
			jobs[i].From = jobs[i-1].To + 1
		}
		if i < d.totalPart-1 {
			jobs[i].To = jobs[i].From + eachSize
		} else {
			// 最后一个分片
			jobs[i].To = fileTotalSize - 1
		}
	}

	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		go func(job filePart) {
			defer wg.Done()
			err := d.downloadPart(job)
			if err != nil {
				log.Errorf("下载文件失败, Error: %s, Job: %v", err, job)
			}
		}(j)
	}
	wg.Wait()

	return d.mergeFileParts()
}

/*
	获取要下载的文件的响应头(header)基本信息
	使用HTTP Method Head方法
*/
func (d *fileDownload) getHeaderInfo() (int, error) {
	headers := map[string]string{
		"User-Agent": userAgent,
	}
	res, err := getNewRequest(d.url, http.MethodHead, headers)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("请求失败, HTTP Code: %v", resp.StatusCode))
	}
	// 检查是否支持断点续传
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("服务器不支持文件断点续传")
	}

	// 支持文件断点续传时，获取文件大小，名称等信息
	outFileName, err := parseFileInfo(resp)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("获取文件信息失败, Error: %v", err))
	}
	if d.outFileName == "" {
		d.outFileName = outFileName
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

// 下载分片
func (d *fileDownload) downloadPart(c filePart) (err error) {
	headers := map[string]string{
		"User-Agent": userAgent,
		"Range":      fmt.Sprintf("bytes=%v-%v", c.From, c.To),
	}
	res, err := getNewRequest(d.url, http.MethodGet, headers)
	if err != nil {
		return
	}
	log.Infof("开始[%d]下载 from: %d to: %d\n", c.Index, c.From, c.To)
	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		return
	}
	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("请求失败, HTTP Code: %v", resp.StatusCode))
	}
	defer resp.Body.Close()
	byteRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(byteRes) != (c.To - c.From + 1) {
		return errors.New("下载错误: 文件分片长度错误")
	}

	c.Data = byteRes
	d.doneFilePart[c.Index] = c
	return
}

// 合并下载的文件
func (d *fileDownload) mergeFileParts() (err error) {
	log.Infof("开始合并文件")
	if !FileExists(d.outFileName) {
		_, err := CreatNestedFile(d.outFileName)
		if err != nil {
			return err
		}
	}

	fileMd5 := sha256.New()
	totalSize := 0
	for _, s := range d.doneFilePart {
		err := ioutil.WriteFile(d.outFileName, s.Data, 0777)
		if err != nil {
			return err
		}
		fileMd5.Write(s.Data)
		totalSize += len(s.Data)
	}
	if totalSize != d.fileSize {
		return errors.New("文件不完整")
	}

	if d.md5 != "" {
		if hex.EncodeToString(fileMd5.Sum(nil)) != d.md5 {
			return errors.New("文件损坏")
		} else {
			log.Info("文件校验成功")
		}
	}
	return
}

func getNewRequest(url, method string, headers map[string]string) (*http.Request, error) {
	r, err := http.NewRequest(
		method,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return r, nil
}

func parseFileInfo(resp *http.Response) (string, error) {
	contentDisposition := resp.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err != nil {
			return "", err
		}
		return params["filename"], nil
	}

	filename := filepath.Base(resp.Request.URL.Path)
	return filename, nil
}
