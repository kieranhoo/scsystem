package routes

import (
	"scsystem/internal/api/controller"
	"scsystem/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Room(a *fiber.App) {
	r := a.Group("/room")
	r.Get("/", controller.Room)
	r.Post("/", controller.NewRoom)
	r.Post("/register", middleware.RegisterRoomValidate, controller.RegisterRoom)
	r.Get("/history", middleware.DatetimeValidate, controller.Histories)

	r.Post("/activity", middleware.ValidateCheckIn, controller.SaveActivityType)
	r.Get("/activity", controller.GetActivity)
}
