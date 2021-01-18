package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// BookTrip ...
func BookTrip(ctx workflow.Context, e BookedEvent) error {
	logger := workflow.GetLogger(ctx)
	defer logger.Info("BookTrip: Workflow completed.")

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    5 * time.Second,
			MaximumAttempts:    5,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	ctx = workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{})

	err := workflow.ExecuteActivity(ctx, BookHotel, e).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, ReserveCar, e).Get(ctx, nil)
	if err != nil {
		cerr := workflow.ExecuteActivity(ctx, CancelHotel, e).Get(ctx, nil)
		if cerr != nil {
			return cerr
		}

		return err
	}

	err = workflow.ExecuteActivity(ctx, BookFlight, e).Get(ctx, nil)
	if err != nil {
		cerr := workflow.ExecuteActivity(ctx, CancelCar, e).Get(ctx, nil)
		if cerr != nil {
			return cerr
		}

		cerr = workflow.ExecuteActivity(ctx, CancelHotel, e).Get(ctx, nil)
		if cerr != nil {
			return cerr
		}

		return err
	}

	return nil
}
