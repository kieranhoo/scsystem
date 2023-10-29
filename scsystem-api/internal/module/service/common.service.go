package service

import (
	"qrcheckin/internal/module/tasks"
	"qrcheckin/pkg/x/mailers"
	"qrcheckin/pkg/x/worker"
)

func Ping() {
	// mailers.SendHTML("iduchungho@gmail.com")
	// log.Println("PONG")
}

func HealthCheck() error {
	return worker.Execute("default_queue", "Worker.HealthCheck", tasks.HealthCheck, int64(1))
}

func Email() (string, error) {
	err := mailers.SendHTML("pkg/template/email_test.html", "iduchungho@gmail.com")
	return "", err
}
