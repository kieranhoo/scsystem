package tasks

import (
	"scsystem/internal/config"
	"scsystem/pkg/x/worker"
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
	WorkerHealthCheck      string = "Worker.HealthCheck"
	WorkerSignUp           string = "Worker.SignUp"
	WorkerSaveUser         string = "Worker.SaveUser"
	WorkerSaveRegistration string = "Worker.SaveRegistration"
	WorkerSaveActivityType string = "Worker.SaveActivityType"
)

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck:      HandleHealthCheck,
		WorkerSaveActivityType: HandleSaveActivityType,
		WorkerSaveRegistration: HandleSaveRegistration,
		WorkerSaveUser:         HandleSaveUser,
	}
}

func Queue() worker.Queue {
	return worker.Queue{
		CriticalQueue: 6, // processed 60% of the time
		DefaultQueue:  3, // processed 30% of the time
		LowQueue:      1, // processed 10% of the time
	}
}
