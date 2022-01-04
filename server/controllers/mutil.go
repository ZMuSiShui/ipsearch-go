package controllers

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/ZMuSiShui/ipsearch-go/conf"
	"github.com/ZMuSiShui/ipsearch-go/util"
	"github.com/gofiber/fiber/v2"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/oschwald/geoip2-golang"
)

var validIPDBs = []string{"ipip", "maxmind", "IP2Location"}

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
	} else if data.IPDB == "IP2Location" {
		ipinfo, err = searchIPbyIP2L(iplist)
	} else {
		ipinfo, err = searchIPbyMaxmind(iplist)
	}
	return
}

// 从 IPIP 数据库查询
func searchIPbyIPIP(ipdata []string) (ipinfolist string, err error) {
	db, err := ipdb.NewCity(conf.IPIPFile)
	if err != nil {
		return
	}
	for _, i := range ipdata {
		i = strings.TrimSpace(i)
		iplist := strings.Split(i, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfolist = ipinfolist + fmt.Sprintf("%s 该 IP 格式不正确\n", i)
		} else {
			ipdata, err := db.FindInfo(iplist[0], "CN")
			if err != nil {
				return "", err
			}
			ipinfo := fmt.Sprintf("%s %s %s %s %s %s %s", i, ipdata.CountryName, ipdata.RegionName, ipdata.CityName, ipdata.OwnerDomain, ipdata.IspDomain, ipdata.CountryCode)
			reg := regexp.MustCompile(`\\s+`)
			ipinfolist += reg.ReplaceAllString(ipinfo, " ")
			ipinfolist = ipinfolist + "\n"
		}
	}
	return
}

// 从 Maxmind 数据库查询
func searchIPbyMaxmind(ipdata []string) (ipinfolist string, err error) {
	db, err := geoip2.Open(conf.MaxmindFile)
	if err != nil {
		return
	}
	defer db.Close()
	for _, i := range ipdata {
		i = strings.TrimSpace(i)
		iplist := strings.Split(i, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfolist = ipinfolist + fmt.Sprintf("%s 该 IP 格式不正确\n", i)
		} else {
			ip := net.ParseIP(iplist[0])
			record, err := db.City(ip)
			if err != nil {
				return "", err
			}
			var countryName string
			var cityName string
			if record.Country.Names["zh-CN"] != "" {
				countryName = record.Country.Names["zh-CN"]
			} else {
				countryName = record.Country.Names["en"]
			}
			if record.City.Names["zh-CN"] != "" {
				cityName = record.City.Names["zh-CN"]
			} else {
				cityName = record.City.Names["en"]
			}
			ipinfo := fmt.Sprintf("%s %s %s %s", i, countryName, cityName, record.Location.TimeZone)
			ipinfo = strings.Replace(ipinfo, "None", "", -1)
			reg := regexp.MustCompile(`\\s+`)
			ipinfolist += reg.ReplaceAllString(ipinfo, " ")
			ipinfolist = ipinfolist + "\n"
		}
	}

	return
}

// // 从 CZ88 数据库查询
// func searchIPbyCZ88(ipdata []string) (ipinfolist string, err error) {
// 	IPDict := util.NewIPDict()
// 	err = IPDict.Load(conf.CZ88File)
// 	if err != nil {
// 		return
// 	}
// 	for _, i := range ipdata {
// 		i = strings.TrimSpace(i)
// 		iplist := strings.Split(i, "/")
// 		ipAddress := net.ParseIP(iplist[0])
// 		if ipAddress == nil {
// 			ipinfolist = ipinfolist + fmt.Sprintf("%s 该 IP 格式不正确\n", i)
// 		} else {
// 			res, err := IPDict.FindIP(iplist[0])
// 			if err != nil {
// 				return "", err
// 			}
// 			ipinfo := fmt.Sprintf("%s %s %s", i, res.Country, res.Area)
// 			reg := regexp.MustCompile(`\\s+`)
// 			ipinfolist += reg.ReplaceAllString(ipinfo, " ")
// 			ipinfolist = ipinfolist + "\n"
// 		}

// 	}
// 	return
// }

func searchIPbyIP2L(ipdata []string) (ipinfolist string, err error) {
	db, err := ip2location.OpenDB(conf.IP2LocationFile)
	if err != nil {
		return
	}
	defer db.Close()
	for _, i := range ipdata {
		i = strings.TrimSpace(i)
		iplist := strings.Split(i, "/")
		if iplist[0] == "" {
			continue
		}
		ipAddress := net.ParseIP(iplist[0])
		if ipAddress == nil {
			ipinfolist = ipinfolist + fmt.Sprintf("%s 该 IP 格式不正确\n", i)
		} else {
			results, err := db.Get_all(iplist[0])
			if err != nil {
				return "", err
			}
			countryName := results.Country_long
			regionName := results.Region
			cityName := results.City
			countryNameCode := results.Country_short
			ipinfo := fmt.Sprintf("%s %s %s %s %s", i, countryName, regionName, cityName, countryNameCode)
			ipinfo = strings.Replace(ipinfo, "None", "", -1)
			reg := regexp.MustCompile(`\\s+`)
			ipinfolist += reg.ReplaceAllString(ipinfo, " ")
			ipinfolist += "\n"
		}
	}
	return
}
