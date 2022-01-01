package controllers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/matryer/try"
)

type respData struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
}

var validwebdb string = "web"

func SingelSearch(c *fiber.Ctx) error {
	const ipApiBaseURL string = "http://ip-api.com/json/%s?lang=zh-CN"
	var search searchReq
	if err := c.BodyParser(&search); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "miss data",
		})
	}
	if search.IPDB != "" && search.IPDB != validwebdb {
		err := fmt.Sprintf("Error: IPDB must be : %s.", validwebdb)
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}
	ipdata := strings.TrimSpace(search.IPData)
	iplist := strings.Split(ipdata, "/")
	ipAddress := net.ParseIP(iplist[0])
	if ipAddress == nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "该 IP 格式不正确",
		})
	}

	data, err := searchIPbyIPAPI(ipApiBaseURL, iplist[0])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "查询失败",
		})
	}
	var msgData map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &msgData)
	return c.JSON(fiber.Map{
		"code": 20000,
		"msg":  "success",
		"data": msgData,
	})
}

func searchIPbyIPAPI(ipapiurl string, ip string) (ipApidoc respData, err error) {
	var request *http.Request
	searchURL := fmt.Sprintf(ipapiurl, ip)
	request, err = http.NewRequest(http.MethodGet, searchURL, nil)
	if err != nil {
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 3,
			DisableKeepAlives:   false,
		},
		Timeout: time.Duration(6) * time.Second,
	}
	var resp *http.Response

	err = try.Do(func(attempt int) (bool, error) {
		var rErr error
		resp, rErr = client.Do(request)
		return attempt < 3, rErr
	})

	if err != nil {
		return
	}
	defer resp.Body.Close()

	var syncRespBodyBytes []byte
	syncRespBodyBytes, err = getResponseBody(resp)
	if err != nil {
		return
	}
	err = json.Unmarshal(syncRespBodyBytes, &ipApidoc)
	return
}

func getResponseBody(resp *http.Response) (body []byte, err error) {
	var output io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		output, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
	default:
		output = resp.Body
		if err != nil {
			return
		}
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(output)
	if err != nil {
		return
	}
	body = buf.Bytes()
	return
}
