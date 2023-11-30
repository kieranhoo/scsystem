package test

import (
	"fmt"
	"log"
	"scsystem/internal/repo"
	"scsystem/internal/service"
	"scsystem/internal/tasks"
	"testing"
)

func TestHistory(t *testing.T) {
	his, err := service.NewLab().GetHistories("100")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(his)
}

func TestUpdate(t *testing.T) {
	log.Println(tasks.SaveRegistration("7h-9h", "DHD", "2013-08-05 18:19:03", "2013-08-05 18:19:03", "12345678", "1"))
}

func TestUser(t *testing.T) {
	log.Println(new(repo.Users).GetByID("2013384"))
}
