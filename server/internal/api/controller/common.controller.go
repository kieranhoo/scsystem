package controller

import (
	"scsystem/internal/service"
	"scsystem/internal/tasks"
	"scsystem/pkg/sentry"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/healthcheck [GET]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func HealthCheckWorker(c *fiber.Ctx) error {
	go tasks.HealthCheck()
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func Email(c *fiber.Ctx) error {
	_, err := service.Email()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func Sentry(c *fiber.Ctx) error {
	panic("y tho")
}

func Capture(c *fiber.Ctx) error {
	sentry.CaptureMessage(c, "User provided unwanted query string, but we recovered just fine")
	return c.SendStatus(fiber.StatusOK)
}
