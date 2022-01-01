package controllers

import "github.com/gofiber/fiber/v2"

func SingelSearch(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": 20000,
		"msg":  "success",
		"data": "data",
	})
}
