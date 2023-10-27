package tasks

import (
	"context"
	"encoding/json"
	"qrcheckin/internal/module/model"
	"qrcheckin/internal/module/types"

	"github.com/hibiken/asynq"
)

func SaveRegistration(registrationTime, supervisor, startDay, endDay, userId, roomId string) error {
	// regis, err := new(model.Registration).GetByUserIdAndRoom(userId, roomId)
	// if err != nil {
	// 	return err
	// }
	// if regis.Empty() {
	// 	return regis.Insert(&model.Registration{
	// 		RegistrationTime: registrationTime,
	// 		Supervisor:       supervisor,
	// 		RoomId:           roomId,
	// 		UserID:           userId,
	// 		StartDay:         startDay,
	// 		EndDay:           endDay,
	// 	})
	// }

	// return regis.UpdateByIDAndRoom(&model.Registration{
	// 	Id:               regis.Id,
	// 	RegistrationTime: registrationTime,
	// 	Supervisor:       supervisor,
	// 	RoomId:           regis.RoomId,
	// 	UserID:           regis.UserID,
	// 	StartDay:         startDay,
	// 	EndDay:           endDay,
	// })

	return model.NewRegistration().Insert(&types.Registration{
		RegistrationTime: registrationTime,
		Supervisor:       supervisor,
		RoomId:           roomId,
		UserID:           userId,
		StartDay:         startDay,
		EndDay:           endDay,
	})
}

func HandleSaveRegistration(_ context.Context, t *asynq.Task) error {
	var registration types.Registration
	if err := json.Unmarshal(t.Payload(), &registration); err != nil {
		return err
	}
	return model.NewRegistration().Insert(&registration)
}
