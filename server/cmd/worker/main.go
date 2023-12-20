package main

import (
	"log"
	"scsystem/api"
)

func main() {
	w := api.NewWorker(10)
	if err := w.Run(); err != nil {
		log.Fatal(err)
	}
}
