package main

import (
	"log"
	"scsystem/internal/messaging"
)

func main() {
	worker := messaging.NewMessaging()
	if err := worker.Run(10); err != nil {
		log.Fatal(err)
	}
}
