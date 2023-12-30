package controller

import (
	"scsystem/internal/schema"
	"scsystem/internal/service"
	"scsystem/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoom
// @Description Register Room.
// @Tags room
// @Accept json
// @Produce json
// @Param sign_in body schema.RegistrationRoomRequest true "RegisterRoom"
// @Success 200 {object} schema.Response
// @Router /room/register [POST]
func RegisterRoom(c *fiber.Ctx) error {
	var RoomService = service.NewRoom()
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

// GetActivity
// @Description Get User Activity for checkin.
// @Tags room
// @Accept json
// @Produce json
// @Param sid query string true "student id"
// @Param room query string true "room id"
// @Success 200 {object} schema.DataResponse
// @Router /room/activity [GET]
func GetActivity(c *fiber.Ctx) error {
	var RoomService = service.NewRoom()
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
// @Param room_id query string true "room id"
// @Param date query string true "date only"
// @Success 200 {object} schema.HistoryDataResponse
// @Router /room/history [GET]
func Histories(c *fiber.Ctx) error {
	var RoomService = service.NewRoom()
	room_id := c.Query("room_id", "")
	if room_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params room_id",
		})
	}
	date := c.Query("date", "")
	if date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params date",
		})
	}
	data, err := RoomService.GetHistoriesData(date, room_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(data)
}

// SaveActivityType
// @Description Save activity type in/out.
// @Tags room
// @Accept json
// @Produce json
// @Param activity_param body schema.CheckInRequest true "CheckInRequest"
// @Success 200 {object} schema.Response
// @Router /room/activity [POST]
func SaveActivityType(c *fiber.Ctx) error {
	var RoomService = service.NewRoom()
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
// @Router /room [GET]
func Room(c *fiber.Ctx) error {
	var RoomService = service.NewRoom()
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
