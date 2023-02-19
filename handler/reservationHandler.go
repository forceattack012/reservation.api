package handler

import (
	"net/http"
	"strconv"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/gofiber/fiber/v2"
)

type ReservationHandler struct {
	service domain.ReservationService
}

func NewReservationHandler(service domain.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (r *ReservationHandler) GetReservationByRoomId(c *fiber.Ctx) error {
	id := c.AllParams()["roomId"]
	roomId, _ := strconv.Atoi(id)
	response, err := r.service.GetReservationByRoomId(roomId)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (r *ReservationHandler) CreateReservation(c *fiber.Ctx) error {
	var reservation domain.Reservation
	var err error
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	err = r.service.CreateReservation(&reservation)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(map[string]string{
		"ID": strconv.Itoa(int(reservation.Id)),
	})
}

func (h *ReservationHandler) DeleteReservation(c *fiber.Ctx) error {
	strId := c.AllParams()["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.service.DeleteReservation(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
