package handler

import (
	"net/http"
	"strconv"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/gofiber/fiber/v2"
)

type PersonHandler struct {
	service domain.PersonService
}

func NewPersonHandler(service domain.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) GetAll(c *fiber.Ctx) error {
	var personList []domain.Person
	personList, err := h.service.GetAllPerson()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(personList)
}

func (h *PersonHandler) NewPerson(c *fiber.Ctx) error {
	person := new(domain.Person)
	var err error
	if err = c.BodyParser(person); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err = h.service.CreatePerson(person); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(map[string]interface{}{
		"ID": person.Id,
	})
}

func (h *PersonHandler) DeletePerson(c *fiber.Ctx) error {
	strId := c.AllParams()["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.service.DeletePersonById(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
