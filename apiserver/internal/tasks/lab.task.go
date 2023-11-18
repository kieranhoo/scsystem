package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"qrcheckin/internal/model"
	"qrcheckin/internal/repo"
)

func SaveActivityType(registrationId, adminId, activityType string) error {
	_history, err := repo.NewHistory().Latest(registrationId)
	if err != nil {
		return err
	}
	if _history.ActivityType == activityType {
		return nil
	}
	return new(repo.History).Insert(&model.History{
		RegistrationId: registrationId,
		ActivityType:   activityType,
		AdminId:        adminId,
	})
}

func HandleSaveActivityType(_ context.Context, task *asynq.Task) error {
	var history model.History
	if err := json.Unmarshal(task.Payload(), &history); err != nil {
		return err
	}
	// fmt.Println(history)
	return repo.NewHistory().Insert(&history)
}
