package tasks

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"qrcheckin/internal/model"
	"qrcheckin/internal/repo"
)

func SaveRegistration(registrationTime, supervisor, startDay, endDay, userId, roomId string) error {
	return repo.NewRegistration().Insert(&model.Registration{
		RegistrationTime: registrationTime,
		Supervisor:       supervisor,
		RoomId:           roomId,
		UserID:           userId,
		StartDay:         startDay,
		EndDay:           endDay,
	})
}

func HandleSaveRegistration(_ context.Context, t *asynq.Task) error {
	var registration model.Registration
	if err := json.Unmarshal(t.Payload(), &registration); err != nil {
		return err
	}
	return repo.NewRegistration().Insert(&registration)
}
