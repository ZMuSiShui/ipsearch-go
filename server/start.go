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

	// 压缩静态资源未gzip或br
	app.Use(compress.New())
	app.Use(cors.New())
	// 初始化路由
	Router(app)

	// 启动
	log.Fatal(app.Listen("0.0.0.0:3000"))
}
