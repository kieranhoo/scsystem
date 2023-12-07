package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func HealthCheck(input int64) error {
	fmt.Printf("HealthCheck Number: %d\n", input)
	return nil
}

func HandleHealthCheck(_ context.Context, task *asynq.Task) error {
	var num int64
	if err := json.Unmarshal(task.Payload(), &num); err != nil {
		return err
	}
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
