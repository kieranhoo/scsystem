// Copyright 2023 Duc Hung Ho. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.
package main

import (
	"os"
	"os/signal"
	"qrcheckin/api"
	middleware2 "qrcheckin/api/middleware"
	routes2 "qrcheckin/api/routes"
	"qrcheckin/cmd/app"
	_ "qrcheckin/docs"
	"qrcheckin/internal/config"
	"qrcheckin/internal/tasks"
	"qrcheckin/pkg/sentry"
	"qrcheckin/pkg/x/mailers"
	"qrcheckin/pkg/x/worker"
	"syscall"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

// @title Student Checkin System
// @version 1.0.0
// @description This is a documentation for the Student Checkin System API
// @host
// @basePath /
func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	coreAPI := api.New().Shutdown(sigChan)
	coreAPI.BackgroundTask(app.JobLaunch)
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
