package main

import (
	"log"
	"os"
	"os/signal"
	"scsystem/api"
	"syscall"
)

// @title Student Checkin System
// @version 1.0.0
// @description This is a documentation for the Student Checkin System API
// @host
// @basePath /
func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	coreAPI := api.NewServer().Shutdown(sigChan)
	if err := coreAPI.Run(); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}
