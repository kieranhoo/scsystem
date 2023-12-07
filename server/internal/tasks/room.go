package tasks

import (
	"context"
	"encoding/json"
	"scsystem/internal/model"
	"scsystem/internal/repo"

	"github.com/hibiken/asynq"
)

func SaveActivityType(history *model.History) error {
	_history, err := repo.NewHistory().Latest(history.RegistrationId)
	if err != nil {
		return err
	}
	if _history.ActivityType == history.ActivityType {
		return nil
	}
	return repo.NewHistory().Insert(history)
}

func HandleSaveActivityType(_ context.Context, task *asynq.Task) error {
	var history model.History
	if err := json.Unmarshal(task.Payload(), &history); err != nil {
		return err
	}
	return SaveActivityType(&history)
}

