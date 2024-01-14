package middleware

import (
	"scsystem/internal/domain"
	"scsystem/internal/misc/guard"
	"scsystem/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func DatetimeValidate(c *fiber.Ctx) error {
	date := c.Query("date", "")
	if date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     "missing query params date",
		})
	}
	if err := guard.Date(date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     "invalid date time format",
		})
	}
	return c.Next()
}

func RegisterRoomValidate(c *fiber.Ctx) error {
	var req domain.RegistrationRoomRequest
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
	if err := guard.Date(req.StartDay); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     "invalid date time format",
		})
	}
	if err := guard.Date(req.EndDay); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Error{
			Success: false,
			Msg:     "invalid date time format",
		})
	}
	return c.Next()
}
