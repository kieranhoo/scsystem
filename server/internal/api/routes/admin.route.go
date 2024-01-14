package routes

import (
	"scsystem/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func Admin(a *fiber.App) {
	r := a.Group("/admin")

	r.Get("/offices", controller.GetAllOffice)
	r.Post("/offices", controller.NewOffice)

	r.Get("/orgs", controller.GetAllOrganizations)
	r.Post("/orgs", controller.NewOrganization)
}
