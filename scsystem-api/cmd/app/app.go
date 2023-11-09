package app

import (
	"os"
	"os/signal"
	"qrcheckin/api"
	middleware2 "qrcheckin/api/middleware"
	routes2 "qrcheckin/api/routes"
	"qrcheckin/internal/config"
	"qrcheckin/internal/service"
	"qrcheckin/internal/tasks"
	"qrcheckin/pkg/x/job"
	"qrcheckin/pkg/x/worker"
	"syscall"
	"time"
)

// Server
// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @basePath /
func Server() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	coreAPI := api.New().Shutdown(sigChan)

	coreAPI.BackgroundTask(JobLaunch)
	coreAPI.Worker(tasks.Setting(config.BrokerUrl, config.ResultBackend))
	coreAPI.Middleware(
		middleware2.FiberMiddleware,
		middleware2.SentryMiddleware,
	)

	coreAPI.Route(
		routes2.Gateway,
		routes2.NotFoundRoute,
	)

	coreAPI.Run()
}

// WorkerLaunch
// Deprecated
func WorkerLaunch(queueName, consume string, concurrency int) error {
	wcf := tasks.Setting(config.BrokerUrl, config.ResultBackend)
	cnf := worker.NewWorker(queueName, wcf)
	return worker.Launch(cnf, consume, concurrency)
}

func AsyncWorker(concurrency int) error {
	w := worker.NewServer(concurrency, worker.Queue{
		config.CriticalQueue: 6, // processed 60% of the time
		config.DefaultQueue:  3, // processed 30% of the time
		config.LowQueue:      1, // processed 10% of the time
	})
	w.HandleFunctions(tasks.Path())
	return w.Run()
}

func JobLaunch() {
	j := job.New()
	if err := j.Scheduler(service.Ping, 5*time.Second); err != nil {
		panic(err)
	}
	if err := j.Launch(); err != nil {
		panic(err)
	}
}
