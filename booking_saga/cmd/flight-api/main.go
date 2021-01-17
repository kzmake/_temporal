package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/seehuhn/mt19937"
)

var r = rand.New(mt19937.New())

func main() {
	r.Seed(time.Now().UnixNano())

	app := fiber.New()

	app.Use(logger.New())
	app.All("/book", bookFlight)
	app.All("/cancel", cancelFlight)

	app.Listen(":3003")
}

func bookFlight(c *fiber.Ctx) error {
	// ko: 80%
	if r.Intn(100) < 80 {
		return fiber.ErrInternalServerError
	}

	// ok: 20%
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"status": "ok"})
}

func cancelFlight(c *fiber.Ctx) error {
	// ko: 0%
	if r.Intn(100) < 0 {
		return fiber.ErrInternalServerError
	}

	// ok: 100%
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"status": "ok"})
}
