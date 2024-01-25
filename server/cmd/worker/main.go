package main

import (
	"log"
	broker "scsystem/internal/broker"
)

func main() {
	worker := broker.Newbroker()
	if err := worker.Run(10); err != nil {
		log.Fatal(err)
	}
}
