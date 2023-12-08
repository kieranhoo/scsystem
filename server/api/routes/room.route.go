package routes

import (
	"scsystem/api/controller"

	"github.com/gofiber/fiber/v2"
)

func RoomRoutes(a *fiber.App) {
	r := a.Group("/v1/room")
	r.Get("/", controller.Room)
	r.Post("/register", controller.RegisterRoom)
	r.Get("/activity", controller.GetActivity)
	r.Get("/history", controller.Histories)
	r.Post("/activity", controller.SaveActivityType)
}
