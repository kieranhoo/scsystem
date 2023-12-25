package service

import (
	"scsystem/internal/tasks"
	"scsystem/pkg/mailers"
	"scsystem/pkg/worker"
)


func HealthCheck() error {
	return worker.Exec(
		tasks.CriticalQueue,
		worker.NewTask(
			tasks.WorkerHealthCheck,
			1,
		),
	)
}

func Email() (string, error) {
	err := mailers.SendHTML("pkg/template/email_test.html", "iduchungho@gmail.com")
	return "", err
}
