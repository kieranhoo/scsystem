package tasks

import (
	"context"
	"encoding/json"
	"scsystem/internal/model"
	"scsystem/internal/repo"

	"github.com/hibiken/asynq"
)

func SaveRegistration(registration *model.Registration) error {
	return repo.NewRegistration().Insert(registration)
}

func HandleSaveRegistration(_ context.Context, t *asynq.Task) error {
	var registration model.Registration
	if err := json.Unmarshal(t.Payload(), &registration); err != nil {
		return err
	}
	return SaveRegistration(&registration)
}
