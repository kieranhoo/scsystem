package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"qrcheckin/internal/module/entity"
	"qrcheckin/internal/module/model"
)

func SaveRegistration(registrationTime, supervisor, startDay, endDay, userId, roomId string) error {
	return model.NewRegistration().Insert(&entity.Registration{
		RegistrationTime: registrationTime,
		Supervisor:       supervisor,
		RoomId:           roomId,
		UserID:           userId,
		StartDay:         startDay,
		EndDay:           endDay,
	})
}

func HandleSaveRegistration(_ context.Context, t *asynq.Task) error {
	var registration entity.Registration
	if err := json.Unmarshal(t.Payload(), &registration); err != nil {
		return err
	}
	return model.NewRegistration().Insert(&registration)
}
