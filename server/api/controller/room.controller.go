package controller

import (
	"scsystem/internal/schema"
	"scsystem/internal/service"
	"scsystem/pkg/x/validator"

	"github.com/gofiber/fiber/v2"
)

var RoomService = service.NewRoom()

// RegisterRoom
// @Description Register Room.
// @Tags room
// @Accept json
// @Produce json
// @Param sign_in body schema.RegistrationRoomRequest true "RegisterRoom"
// @Success 200 {object} schema.Response
// @Router /v1/room/register [POST]
func RegisterRoom(c *fiber.Ctx) error {
	var req schema.RegistrationRoomRequest
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
		if err := RoomService.RegisterRoom(&req); err != nil {
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
// @Tags room
// @Accept json
// @Produce json
// @Param sid query string true "student id"
// @Param room query string true "room id"
// @Success 200 {object} schema.DataResponse
// @Router /v1/room/user [GET]
func GetActivity(c *fiber.Ctx) error {
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
	_RoomUser, err := RoomService.RegistrationLatest(sid, room)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	if _RoomUser.RegistrationId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "no data available",
		})
	}
	return c.JSON(schema.DataResponse{
		Success: true,
		Data:    _RoomUser,
	})
}

// Histories
// @Description Get History.
// @Tags room
// @Accept json
// @Produce json
// @Param limit query string true "limit records"
// @Success 200 {object} schema.DataResponse
// @Router /v1/room/history [GET]
func Histories(c *fiber.Ctx) error {
	limit := c.Query("limit", "")
	if limit == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params limit",
		})
	}
	data, err := RoomService.GetHistoriesData(limit)
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
// @Tags room
// @Accept json
// @Produce json
// @Param activity_param body schema.CheckInRequest true "CheckInRequest"
// @Success 200 {object} schema.Response
// @Router /v1/room/activity [POST]
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
		if err := RoomService.SaveActivityType(&req); err != nil {
			panic(err.Error())
		}
	}()
	return c.JSON(schema.Response{
		Success: true,
		Msg:     "successfully",
	})
}

// Room
// @Description Get Room information from database
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {object} schema.DataResponse
// @Router /v1/room [GET]
func Room(c *fiber.Ctx) error {
	data, err := RoomService.GetRoom()
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
