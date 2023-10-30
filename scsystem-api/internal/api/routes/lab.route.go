package routes

import (
	"qrcheckin/internal/api/controller"
	"qrcheckin/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func LabRoutes(a *fiber.App) {
	r := a.Group("/api/v1/lab")
	r.Post("/register", controller.RegisterLab)
	r.Get("/user", controller.GetUser)
	r.Get("/history", controller.Histories)
	r.Post("/activity", middleware.ValidateCheckIn, controller.SaveActivityType)
}
