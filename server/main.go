// Copyright 2023 Duc Hung Ho. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.
package main

import (
	"log"
	"os"
	"os/signal"
	_ "scsystem/docs"
	"scsystem/internal/api"
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
