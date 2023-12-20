package routes

import (
	"scsystem/api/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	r := a.Group("/auth")
	r.Post("/signup", controller.SignUp)
	r.Post("/signin", controller.SignIn)
}

func UserRoutes(a *fiber.App) {
	r := a.Group("/user")
	r.Get("/", controller.GetMe)
}
