package queue

import (
	"scsystem/internal/config"
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
