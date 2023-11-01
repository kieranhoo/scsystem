package tasks

import (
	"context"
	"encoding/json"
	"qrcheckin/internal/mod/model"
	"qrcheckin/internal/types/entity"

	"github.com/hibiken/asynq"
)

func SaveActivityType(registrationId, adminId, activityType string) error {
	_history, err := model.NewHistory().Latest(registrationId)
	if err != nil {
		return err
	}
	if _history.ActivityType == activityType {
		return nil
	}
	return new(model.History).Insert(&entity.History{
		RegistrationId: registrationId,
		ActivityType:   activityType,
		AdminId:        adminId,
	})
}

func HandleSaveActivityType(_ context.Context, task *asynq.Task) error {
	var history entity.History
	if err := json.Unmarshal(task.Payload(), &history); err != nil {
		return err
	}
	// fmt.Println(history)
	return model.NewHistory().Insert(&history)
}
