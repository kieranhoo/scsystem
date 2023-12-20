package routes

import "github.com/gofiber/fiber/v2"

func NotificationRoute(a *fiber.App) {
	r := a.Group("/notifications")
	r.Get("/")
}
