package main

import (
	"log"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	temporal "go.temporal.io/sdk/client"

	"github.com/kzmake/_temporal/booking_saga/workflow"
)

var client temporal.Client

func getEnv(key, val string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return val
}

func main() {
	var err error
	client, err = temporal.NewClient(temporal.Options{
		HostPort: getEnv("TEMPORAL_HOST_PORT", "localhost:7233"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.All("/book", bookTrip)

	app.Listen(":3000")
}

func bookTrip(c *fiber.Ctx) error {
	e := workflow.NewBookedEvent(
		uuid.New().String(),
		randomdata.FullName(randomdata.RandomGender),
		randomdata.Country(randomdata.FullCountry),
	)

	options := temporal.StartWorkflowOptions{
		ID:        e.ID,
		TaskQueue: workflow.BookTripTaskQueue,
	}

	_, err := client.ExecuteWorkflow(c.Context(), options, workflow.BookTrip, e)
	if err != nil {
		log.Println("error starting TransferMoney workflow", err)
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusAccepted).JSON(e)
}
