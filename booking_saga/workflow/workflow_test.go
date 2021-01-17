package workflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testcases := []struct {
		in BookedEvent
	}{
		{in: NewBookedEvent("a41ad012-6473-49b7-b657-14d307829483", "hoge", "Hawaii")},
		{in: NewBookedEvent("7f998bca-8f12-421f-93e2-622bb494de67", "fuga", "Tokyo")},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("すべての Activities が1回目で成功する: %s", tt.in.ID), func(t *testing.T) {
			testSuite := &testsuite.WorkflowTestSuite{}
			env := testSuite.NewTestWorkflowEnvironment()

			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(nil)
			env.OnActivity(BookFlight, mock.Anything, tt.in).Return(nil)

			env.ExecuteWorkflow(BookTrip, tt.in)

			require.True(t, env.IsWorkflowCompleted())
			require.True(t, env.AssertExpectations(t))
			require.NoError(t, env.GetWorkflowError())
		})
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("すべての Activities が5回目で成功する: %s", tt.in.ID), func(t *testing.T) {
			testSuite := &testsuite.WorkflowTestSuite{}
			env := testSuite.NewTestWorkflowEnvironment()

			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(5)
			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(5)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(nil)
			env.OnActivity(BookFlight, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(5)
			env.OnActivity(BookFlight, mock.Anything, tt.in).Return(nil)

			env.ExecuteWorkflow(BookTrip, tt.in)

			require.True(t, env.IsWorkflowCompleted())
			require.True(t, env.AssertExpectations(t))
			require.NoError(t, env.GetWorkflowError())
		})
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("BookFlight が6回とも失敗した + Cancel処理が成功した: %s", tt.in.ID), func(t *testing.T) {
			testSuite := &testsuite.WorkflowTestSuite{}
			env := testSuite.NewTestWorkflowEnvironment()

			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(nil)
			env.OnActivity(BookFlight, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(6)
			env.OnActivity(CancelCar, mock.Anything, tt.in).Return(nil)
			env.OnActivity(CancelHotel, mock.Anything, tt.in).Return(nil)

			env.ExecuteWorkflow(BookTrip, tt.in)

			require.True(t, env.IsWorkflowCompleted())
			require.True(t, env.AssertExpectations(t))
			require.Error(t, env.GetWorkflowError())
		})
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("ReserveCar が6回とも失敗した + Cancel処理が成功した: %s", tt.in.ID), func(t *testing.T) {
			testSuite := &testsuite.WorkflowTestSuite{}
			env := testSuite.NewTestWorkflowEnvironment()

			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(6)
			env.OnActivity(CancelHotel, mock.Anything, tt.in).Return(nil)

			env.ExecuteWorkflow(BookTrip, tt.in)

			require.True(t, env.IsWorkflowCompleted())
			require.True(t, env.AssertExpectations(t))
			require.Error(t, env.GetWorkflowError())
		})
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("ReserveCar が6回とも失敗した + Cancel処理が成功した: %s", tt.in.ID), func(t *testing.T) {
			testSuite := &testsuite.WorkflowTestSuite{}
			env := testSuite.NewTestWorkflowEnvironment()

			env.OnActivity(BookHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(ReserveCar, mock.Anything, tt.in).Return(fmt.Errorf("mocked error")).Times(6)
			env.OnActivity(CancelHotel, mock.Anything, tt.in).Return(nil)
			env.OnActivity(CancelCar, mock.Anything, tt.in).Return(nil)

			env.ExecuteWorkflow(BookTrip, tt.in)

			require.True(t, env.IsWorkflowCompleted())
			require.True(t, env.AssertExpectations(t))
			require.Error(t, env.GetWorkflowError())
		})
	}
}
