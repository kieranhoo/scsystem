package routes

import (
	"scsystem/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Stat(a *fiber.App) {
	r := a.Group("/stat")
	r.Get("/chart", controller.GetChartData)
	r.Get("/rooms", controller.GetRoomData)
}
