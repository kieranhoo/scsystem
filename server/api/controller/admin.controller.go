package controller

import (
	"scsystem/internal/schema"
	"scsystem/internal/service"
	"scsystem/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func NewRoom(c *fiber.Ctx) error {
	var req schema.NewRoomRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if err := service.InsertNewRoom(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "Could not create room",
		})
	}
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "your room was created successfully",
	})
}

func NewOrganization(c *fiber.Ctx) error {
	var req schema.NewOrganizationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if err := service.InsertNewOrganization(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "Could not create organization",
		})
	}
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "your organization was created successfully",
	})
}

func GetAllOrganizations(c *fiber.Ctx) error {
	if data, err := service.GetAllOrganization(); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    data,
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(schema.Response{
		Success: false,
		Msg:     "Could not get organization",
	})
}

func NewOffice(c *fiber.Ctx) error {
	var req schema.NewOfficeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _validate := validator.Validator(req); _validate != "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     _validate,
		})
	}
	if err := service.InsertNewOffice(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "Could not create office",
		})
	}
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "your office was created successfully",
	})
}

func GetAllOffice(c *fiber.Ctx) error {
	if data, err := service.GetAllOffice(); err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    data,
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(schema.Response{
		Success: false,
		Msg:     "Could not get office",
	})
}
