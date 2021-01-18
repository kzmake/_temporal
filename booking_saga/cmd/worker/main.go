package main

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/kzmake/_temporal/booking_saga/workflow"
)

func getEnv(key, val string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return val
}

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: getEnv("TEMPORAL_HOST_PORT", "localhost:7233"),
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.BookTripTaskQueue, worker.Options{
		MaxConcurrentActivityExecutionSize: 1,
	})

	w.RegisterWorkflow(workflow.BookTrip)
	w.RegisterActivity(workflow.BookHotel)
	w.RegisterActivity(workflow.CancelHotel)
	w.RegisterActivity(workflow.ReserveCar)
	w.RegisterActivity(workflow.CancelCar)
	w.RegisterActivity(workflow.BookFlight)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
