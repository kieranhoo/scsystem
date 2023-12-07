package test

import (
	"fmt"
	"log"
	"scsystem/internal/repo"
	"scsystem/internal/service"
	"testing"
)

func TestHistory(t *testing.T) {
	his, err := service.NewRoom().GetHistories("100")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(his)
}

func TestUser(t *testing.T) {
	log.Println(new(repo.Users).GetByID("2013384"))
}
