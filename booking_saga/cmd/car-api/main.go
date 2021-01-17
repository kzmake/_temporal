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
	app.All("/reserve", reserveCar)
	app.All("/cancel", cancelCar)

	app.Listen(":3002")
}

func reserveCar(c *fiber.Ctx) error {
	// ko: 30%
	if r.Intn(100) < 30 {
		return fiber.ErrInternalServerError
	}

	// ok: 70%
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"status": "ok"})
}

func cancelCar(c *fiber.Ctx) error {
	// ko: 10%
	if r.Intn(100) < 10 {
		return fiber.ErrInternalServerError
	}

	// ok:90%
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"status": "ok"})
}
