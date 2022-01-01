package controllers

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ZMuSiShui/ipsearch-go/conf"
	"github.com/ZMuSiShui/ipsearch-go/util"
	"github.com/gofiber/fiber/v2"
	"github.com/ipipdotnet/ipdb-go"
)

var validIPDBs = []string{"ipip", "maxmind", "qqzeng", "cz88"}

type searchReq struct {
	IPDB   string `json:"ipdb"`
	IPData string `json:"ipdata"`
}

func MutilSearch(c *fiber.Ctx) error {
	var search searchReq
	if err := c.BodyParser(&search); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "miss data",
		})
	}

	if search.IPDB != "" && !util.IsContain(search.IPDB, validIPDBs) {
		err := fmt.Sprintf("Error: IPDB must be one of: %s.", strings.Join(validIPDBs, ", "))
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}

	data, err := searchIP(search)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": 20000,
		"msg":  "success",
		"data": data,
	})
}

// 分类查询
func searchIP(data searchReq) (ipinfo string, err error) {
	// 处理 ip 字符串
	iplist := strings.Split(data.IPData, "\n")
	if data.IPDB == "ipip" {
		ipinfo, err = searchIPbyIPIP(iplist)
	} else if data.IPDB == "maxmind" {
		ipinfo, err = searchIPbyMaxmind(iplist)
	} else if data.IPDB == "maxmind" {
		ipinfo, err = searchIPbyQQZeng(iplist)
	} else if data.IPDB == "maxmind" {
		ipinfo, err = searchIPbyCZ88(iplist)
	} else {
		ipinfo, err = searchIPbyMaxmind(iplist)
	}
	return
}

// 从 IPIP 数据库查询
func searchIPbyIPIP(ipdata []string) (ipinfolist string, err error) {
	db, eErr := ipdb.NewCity(conf.IPIPFile)
	if eErr != nil {
		return "", eErr
	}
	for _, i := range ipdata {
		i = strings.TrimSpace(i)
		iplist := strings.Split(i, "/")
		ipdata, err := db.FindInfo(iplist[0], "CN")
		if err != nil {
			return "", err
		}
		ipinfo := fmt.Sprintf("%s %s %s %s %s %s %s\n", i, ipdata.CountryName, ipdata.RegionName, ipdata.CityName, ipdata.OwnerDomain, ipdata.IspDomain, ipdata.CountryCode)
		reg := regexp.MustCompile("\\s+")
		ipinfolist += reg.ReplaceAllString(ipinfo, "")
	}
	return
}

// 从 Maxmind 数据库查询
func searchIPbyMaxmind(ipdata []string) (ipinfo string, err error) {
	return
}

// 从 QQZeng 数据库查询
func searchIPbyQQZeng(ipdata []string) (ipinfo string, err error) {
	return
}

// 从 CZ88 数据库查询
func searchIPbyCZ88(ipdata []string) (ipinfo string, err error) {
	return
}
