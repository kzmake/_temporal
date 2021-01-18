package workflow

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"go.temporal.io/sdk/activity"
)

var (
	hotel  = resty.New().SetHostURL(getEnv("HOTEL_ENDPOINT", "http://localhost:3001"))
	car    = resty.New().SetHostURL(getEnv("CAR_ENDPOINT", "http://localhost:3002"))
	flight = resty.New().SetHostURL(getEnv("FLIGHT_ENDPOINT", "http://localhost:3003"))
)

const (
	bookHotelLogMessagePrefix   = "Activity BookHotel"
	cancelHotelLogMessagePrefix = "Activity CancelHotel"
	reserveCarLogMessagePrefix  = "Activity ReserveCar"
	cancelCarLogMessagePrefix   = "Activity CancelCar"
	bookFlightLogMessagePrefix  = "Activity BookFlight"
)

func getEnv(key, val string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return val
}

// BookHotel ...
func BookHotel(ctx context.Context, e BookedEvent) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("%s: started: %+v", bookHotelLogMessagePrefix, e))
	defer logger.Info(fmt.Sprintf("%s: completed", bookHotelLogMessagePrefix))

	res, err := hotel.R().Post("/book")
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("failed: recv res.Status: %s", res.Status())
	}

	logger.Info(fmt.Sprintf("%s: recv res.Status: %s", bookHotelLogMessagePrefix, res.Status()))

	return nil
}

// CancelHotel ...
func CancelHotel(ctx context.Context, e BookedEvent) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("%s: started: %+v", cancelHotelLogMessagePrefix, e))
	defer logger.Info(fmt.Sprintf("%s: completed", cancelHotelLogMessagePrefix))

	res, err := hotel.R().Post("/cancel")
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("failed: recv res.Status: %s", res.Status())
	}

	logger.Info(fmt.Sprintf("%s: recv res.Status: %s", cancelHotelLogMessagePrefix, res.Status()))

	return nil
}

// ReserveCar ...
func ReserveCar(ctx context.Context, e BookedEvent) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("%s: started: %+v", reserveCarLogMessagePrefix, e))
	defer logger.Info(fmt.Sprintf("%s: completed", reserveCarLogMessagePrefix))

	res, err := car.R().Post("/reserve")
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("failed: recv res.Status: %s", res.Status())
	}

	logger.Info(fmt.Sprintf("%s: recv res.Status: %s", reserveCarLogMessagePrefix, res.Status()))

	return nil
}

// CancelCar ...
func CancelCar(ctx context.Context, e BookedEvent) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("%s: started: %+v", cancelCarLogMessagePrefix, e))
	defer logger.Info(fmt.Sprintf("%s: completed", cancelCarLogMessagePrefix))

	res, err := car.R().Post("/cancel")
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("failed: recv res.Status: %s", res.Status())
	}

	logger.Info(fmt.Sprintf("%s: recv res.Status: %s", cancelCarLogMessagePrefix, res.Status()))

	return nil
}

// BookFlight ...
func BookFlight(ctx context.Context, e BookedEvent) error {
	logger := activity.GetLogger(ctx)
	logger.Info(fmt.Sprintf("%s: started: %+v", bookFlightLogMessagePrefix, e))
	defer logger.Info(fmt.Sprintf("%s: completed", bookFlightLogMessagePrefix))

	res, err := flight.R().Post("/book")
	if err != nil {
		return err
	}

	if res.IsError() {
		return fmt.Errorf("failed: recv res.Status: %s", res.Status())
	}

	logger.Info(fmt.Sprintf("%s: recv res.Status: %s", bookFlightLogMessagePrefix, res.Status()))

	return nil
}
