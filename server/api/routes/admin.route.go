package routes

import "github.com/gofiber/fiber/v2"

func Admin(a *fiber.App) {
	r := a.Group("/admin")
	r.Get("/rooms")
	r.Post("/rooms")

	r.Get("/offices")
	r.Post("/offices")

	r.Get("/orgs")
	r.Post("/orgs")
}
