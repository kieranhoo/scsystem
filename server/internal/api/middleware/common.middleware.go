package middleware

import (
	"scsystem/internal/domain"
	"scsystem/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func ValidateCheckIn(c *fiber.Ctx) error {
	var req domain.CheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if req.ActivityType != "in" && req.ActivityType != "out" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     "invalid activity type, uses 'in' or 'out' lowercase",
		})
	}
	return c.Next()
}
