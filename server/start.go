/*
 * @Author: iRorikon
 * @Date: 2022-06-22 14:22:01
 * @FilePath: \ipsearch-go\server\start.go
 */
package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nekomi-cn/ipsearch-go/config"
	log "github.com/sirupsen/logrus"
)

func init() {

}

func Start() {
	// 创建实例
	app := fiber.New()
	// app.Use(logger.New())   // 开发模式下使用
	app.Use(compress.New()) // 压缩静态资源为gzip或br
	// app.Use(etag.New())     //一些内容不变的东西，不会重复发送
	// app.Use(cache.New(cache.Config{
	// 	Expiration: 2 * time.Minute,
	// })) // 生产环境 缓存一分钟内的请求结果
	app.Use(cors.New())

	// 初始化路由
	Router(app)
	// app.Get("/dashboard", monitor.New()) // 代码运行监视器，开发环境使用
	// 启动
	// log.Fatal(app.Listen(":8080"))
	log.Fatal(app.Listen(fmt.Sprintf("%s:%v", config.CFG.System.Address, config.CFG.System.Port)))
}
