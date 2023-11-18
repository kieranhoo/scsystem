package app

import (
	"os"
	"os/signal"
	"scsystem/api"
	"syscall"
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
func Server() error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	coreAPI := api.New().Shutdown(sigChan)
	return coreAPI.Run()
}

func AsyncWorker(concurrency int) error {
	w := api.NewWorker(10)
	return w.Run()
}
