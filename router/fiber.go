package router

import (
	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRoute() *FiberRouter {
	r := fiber.New()
	return &FiberRouter{r}
}
