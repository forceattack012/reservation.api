package handler

import (
	"net/http"
	"strconv"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	service domain.RoomService
}

func NewRoomHandler(service domain.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

func (r *RoomHandler) GetRoomAll(c *fiber.Ctx) error {
	rooms, err := r.service.GetAllRoom()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(rooms)
}

func (r *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	var room domain.Room
	var err error
	if err = c.BodyParser(&room); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err = r.service.CreateRoom(&room); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(map[string]string{
		"ID": strconv.Itoa(int(room.Id)),
	})
}

func (r *RoomHandler) UpdateRoom(c *fiber.Ctx) error {
	id := c.AllParams()["roomId"]
	roomId, _ := strconv.Atoi(id)

	var room domain.Room
	err := c.BodyParser(&room)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	err = r.service.UpdateRoom(roomId, &room)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusOK)
}
