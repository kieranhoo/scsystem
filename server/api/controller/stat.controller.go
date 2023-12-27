package controller

import (
	"scsystem/internal/schema"
	"scsystem/internal/service"

	"github.com/gofiber/fiber/v2"
)

// GetChartData
// @Description Get Chart Data IN/OUT.
// @Tags stat
// @Accept json
// @Produce json
// @Param room_id query string true "room id"
// @Success 200 {object} schema.ChartMetadata
// @Router /stat/chart [GET]
func GetChartData(c *fiber.Ctx) error {
	roomId := c.Query("room_id", "")
	if roomId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params room_id",
		})
	}
	data, err := service.NewStat().GetChartData(roomId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(data)
}

// GetRoomData
// @Description Get Room Data IN/OUT.
// @Tags stat
// @Accept json
// @Produce json
// @Param room_id query string true "room id"
// @Success 200 {object} schema.RoomStat
// @Router /stat/rooms [GET]
func GetRoomData(c *fiber.Ctx) error {
	roomId := c.Query("room_id", "")
	if roomId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     "missing query params room_id",
		})
	}
	data, err := service.NewStat().GetRoomData(roomId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Success: false,
			Msg:     err.Error(),
		})
	}
	return c.JSON(data)
}
