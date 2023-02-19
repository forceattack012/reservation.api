package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRoute() *FiberRouter {
	r := fiber.New()

	config := cors.Config{
		AllowOrigins: strings.Join([]string{"http://localhost:4200", "http://localhost:3000"}, ","),
	}

	r.Use(cors.New(config))
	r.Use(logger.New())

	return &FiberRouter{r}
}
