package tasks

import (
	"qrcheckin/pkg/x/worker"
)

const (
	DefaultQueue           string = "default_queue"
	WorkerHealthCheck      string = "Worker.HealthCheck"
	WorkerSignUp           string = "Worker.SignUp"
	WorkerSaveUser         string = "Worker.SaveUser"
	WorkerSaveRegistration string = "Worker.SaveRegistration"
	WorkerSaveActivityType string = "Worker.SaveActivityType"
)

func Setting(broker, resultBackend string) *worker.Config {
	return &worker.Config{
		Broker:        broker,
		ResultBackend: resultBackend,
		Tasks: worker.Tasks{
			DefaultQueue: worker.Task{
				WorkerHealthCheck:      HealthCheck,
				WorkerSignUp:           SignUp,
				WorkerSaveUser:         SaveUser,
				WorkerSaveRegistration: SaveRegistration,
				WorkerSaveActivityType: SaveActivityType,
			},
		},
	}
}

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck:      HandleHealthCheck,
		WorkerSaveActivityType: HandleSaveActivityType,
		WorkerSaveRegistration: HandleSaveRegistration,
		WorkerSaveUser:         HandleSaveUser,
	}
}
