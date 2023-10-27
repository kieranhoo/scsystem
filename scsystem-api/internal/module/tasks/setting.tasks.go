package tasks

import (
	worker2 "qrcheckin/pkg/x/worker"
)

const (
	DefaultQueue           string = "default_queue"
	WorkerHealthCheck      string = "Worker.HealthCheck"
	WorkerSignUp           string = "Worker.SignUp"
	WorkerSaveUser         string = "Worker.SaveUser"
	WorkerSaveRegistration string = "Worker.SaveRegistration"
	WorkerSaveActivityType string = "Worker.SaveActivityType"
)

func Setting(broker, resultBackend string) *worker2.Config {
	return &worker2.Config{
		Broker:        broker,
		ResultBackend: resultBackend,
		Tasks: worker2.Tasks{
			DefaultQueue: worker2.Task{
				WorkerHealthCheck:      HealthCheck,
				WorkerSignUp:           SignUp,
				WorkerSaveUser:         SaveUser,
				WorkerSaveRegistration: SaveRegistration,
				WorkerSaveActivityType: SaveActivityType,
			},
		},
	}
}

func Path() worker2.Path {
	return worker2.Path{
		WorkerHealthCheck:      HandleHealthCheck,
		WorkerSaveActivityType: HandleSaveActivityType,
		WorkerSaveRegistration: HandleSaveRegistration,
		WorkerSaveUser:         HandleSaveUser,
	}
}
