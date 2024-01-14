package tasks

import (
	"scsystem/internal/messaging/queue"
	"scsystem/pkg/worker"
)

func HealthCheck() error {
	return worker.Exec(
		queue.CriticalQueue,
		worker.NewTask(
			queue.WorkerHealthCheck,
			1,
		),
	)
}
