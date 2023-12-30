package middleware

import (
	"scsystem/internal/guard"
	"scsystem/internal/schema"

	"github.com/gofiber/fiber/v2"
)

func DatetimeValidate(c *fiber.Ctx) error {
	date := c.Query("date", "")
	if date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params date",
		})
	}
	if err := guard.Date(date); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "invalid date time format",
		})
	}
	return c.Next()
}
