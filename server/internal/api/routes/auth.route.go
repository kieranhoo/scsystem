package routes

import (
	"scsystem/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Auth(a *fiber.App) {
	r := a.Group("/auth")
	r.Post("/signup", controller.SignUp)
	r.Post("/signin", controller.SignIn)
}

func User(a *fiber.App) {
	r := a.Group("/user")
	r.Get("/", controller.GetMe)
}
