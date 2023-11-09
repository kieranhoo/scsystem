package routes

import (
	"github.com/gofiber/fiber/v2"
	"qrcheckin/api/controller"
)

func AuthRoutes(a *fiber.App) {
	r := a.Group("/v1/auth")
	r.Post("/signup", controller.SignUp)
	r.Post("/signin", controller.SignIn)
}
