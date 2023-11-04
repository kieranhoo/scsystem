package controller

import (
	"qrcheckin/internal/mod/schema"
	"qrcheckin/internal/mod/service"
	"qrcheckin/pkg/x/validator"

	"github.com/gofiber/fiber/v2"
)

var labService = service.NewLab()

// RegisterLab
// @Description Register lab.
// @Tags lab
// @Accept json
// @Produce json
// @Param sign_in body schema.RegistrationLabRequest true "RegisterLab"
// @Success 200 {object} schema.Response
// @Router /v1/lab/register [POST]
func RegisterLab(c *fiber.Ctx) error {
	var req schema.RegistrationLabRequest
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
	go func() {
		if err := labService.RegisterLab(&req); err != nil {
			panic(err.Error())
		}
	}()
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "your registration was created successfully",
	})
}

// GetUser
// @Description Get User for checkin.
// @Tags lab
// @Accept json
// @Produce json
// @Param sid query string true "student id"
// @Param room query string true "student id"
// @Success 200 {object} schema.DataResponse
// @Router /v1/lab/user [GET]
func GetUser(c *fiber.Ctx) error {
	sid := c.Query("sid", "")
	if sid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params sid",
		})
	}
	room := c.Query("room", "")
	if room == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params room",
		})
	}
	_labUser, err := labService.RegistrationLatest(sid, room)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _labUser.Id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "no data available",
		})
	}
	return c.JSON(schema.DataResponse{
		Success: true,
		Data:    _labUser,
	})
}

// Histories
// @Description Get History.
// @Tags lab
// @Accept json
// @Produce json
// @Param limit query string true "limit records"
// @Success 200 {object} schema.DataResponse
// @Router /v1/lab/history [GET]
func Histories(c *fiber.Ctx) error {
	limit := c.Query("limit", "")
	if limit == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params limit",
		})
	}
	data, err := labService.GetHistoriesData(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(schema.DataResponse{
		Success: true,
		Data:    data,
	})
}

// SaveActivityType
// @Description Save activity type in/out.
// @Tags lab
// @Accept json
// @Produce json
// @Param limit query string true "limit records"
// @Success 200 {object} schema.Response
// @Router /v1/lab/activity [POST]
func SaveActivityType(c *fiber.Ctx) error {
	var req schema.CheckInRequest
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
	go func() {
		if err := labService.SaveActivityType(&req); err != nil {
			panic(err.Error())
		}
	}()
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "successfully",
	})
}
