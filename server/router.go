package server

import (
	"github.com/ZMuSiShui/ipsearch-go/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// 静态文件目录
	app.Static("/", "web") // 以主程序的路径为根目录
	// api接口
	api := app.Group("/api")
	{
		api.Post("/mutil", controllers.MutilSearch)
		api.Post("/Single", controllers.SingelSearch)
	}
}
