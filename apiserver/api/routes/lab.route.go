package routes

import (
	"github.com/gofiber/fiber/v2"
	"qrcheckin/api/controller"
	"qrcheckin/api/middleware"
)

func LabRoutes(a *fiber.App) {
	r := a.Group("/v1/lab")
	r.Post("/register", controller.RegisterLab)
	r.Get("/user", controller.GetUser)
	r.Get("/history", controller.Histories)
	r.Post("/activity", middleware.ValidateCheckIn, controller.SaveActivityType)
}
