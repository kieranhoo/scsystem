package queue

import (
	"scsystem/internal/config"
	"scsystem/internal/messaging/handler"
	"scsystem/pkg/worker"
)

var (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

func init() {
	if config.StageStatus != "prod" {
		CriticalQueue = "critical_dev"
		DefaultQueue = "default_dev"
		LowQueue = "low_dev"
	}
}

const (
	WorkerHealthCheck string = "Worker.HealthCheck"
)

func Common(eng *worker.Engine) {
	eng.Queue(WorkerHealthCheck, handler.HandleHealthCheck)
}
