package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	// log.Fatal(app.Listen(":8080")) // linux环境
	log.Fatal(app.Listen("127.0.0.1:3000")) // windows环境
}
