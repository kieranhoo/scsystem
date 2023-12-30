package service

import (
	"scsystem/internal/model"
	"scsystem/internal/repo"
	"scsystem/internal/schema"
	"scsystem/internal/tasks"
	"scsystem/pkg/mailers"
	"time"
)

type Registration struct {
	repo repo.IRegistration
}

func NewRoom() IRoomService {
	return &Registration{
		repo: repo.NewRegistration(),
	}
}

func (regis *Registration) RegisterRoom(req *schema.RegistrationRoomRequest) error {
	_user, err := regis.repo.GetByUserIdAndRoom(req.StudentId, req.RoomId)
	// log.Fatal(_user)
	if err != nil || _user.Id == "" {
		if err := tasks.SaveUser(&model.Users{
			Id:          req.StudentId,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
		}); err != nil {
			return err
		}
		// if err := worker.Exec(tasks.CriticalQueue, worker.NewTask(
		// 	tasks.WorkerSaveUser,
		// 	model.Users{
		// 		Id:          req.StudentId,
		// 		FirstName:   req.FirstName,
		// 		LastName:    req.LastName,
		// 		PhoneNumber: req.PhoneNumber,
		// 		Email:       req.Email,
		// 	}),
		// ); err != nil {
		// 	return err
		// }
	}

	go mailers.ConfirmRegistrationRoom(req.Email, mailers.ConfirmEmail{
		Name:        req.FirstName + " " + req.LastName,
		RoomNumber:  req.RoomId,
		Date:        time.Now().UTC().Format(time.DateOnly),
		Time:        req.RegistrationTime,
		StartTime:   req.StartDay,
		EndTime:     req.EndDay,
		Email:       "tarzaines@gmail.com",
		FullName:    "ad",
		Position:    "admin",
		Information: "ist Director",
	})

	return tasks.SaveRegistration(&model.Registration{
		RegistrationTime: req.RegistrationTime,
		Supervisor:       req.Supervisor,
		StartDay:         req.StartDay,
		EndDay:           req.EndDay,
		UserID:           req.StudentId,
		RoomId:           req.RoomId,
	})
	// return worker.Exec(tasks.CriticalQueue, worker.NewTask(
	// 	tasks.WorkerSaveRegistration,
	// 	model.Registration{
	// 		RegistrationTime: req.RegistrationTime,
	// 		Supervisor:       req.Supervisor,
	// 		StartDay:         req.StartDay,
	// 		EndDay:           req.EndDay,
	// 		UserID:           req.StudentId,
	// 		RoomId:           req.RoomId,
	// 	},
	// ))
}

func (regis *Registration) SaveActivityType(req *schema.CheckInRequest) error {
	return tasks.SaveActivityType(&model.History{
		RegistrationId: req.RegistrationId,
		AdminId:        req.AdminId,
		ActivityType:   req.ActivityType,
	})
	// return worker.Exec(tasks.CriticalQueue, worker.NewTask(
	// 	tasks.WorkerSaveActivityType,
	// 	model.History{
	// 		RegistrationId: req.RegistrationId,
	// 		AdminId:        req.AdminId,
	// 		ActivityType:   req.ActivityType,
	// 	},
	// ))
}

func (regis *Registration) RegistrationLatest(studentId, roomId string) (*schema.UserRoomData, error) {
	return regis.repo.RegistrationLatest(studentId, roomId)
}

func (regis *Registration) GetHistories(limit string) ([]model.History, error) {
	return repo.NewHistory().GetList(limit)
}

func (regis *Registration) GetHistoriesData(date, roomId string) (*schema.HistoryDataResponse, error) {
	data, err := repo.NewHistory().GetHistory(date, roomId)
	if err != nil {
		return nil, err
	}
	var response = schema.HistoryDataResponse{
		Data: data,
	}
	if len(data) > 0 {
		response.RoomName = data[0].RoomName
	}
	for _, v := range data {
		switch v.ActivityType {
		case "in":
			response.TotalIn++
		case "out":
			response.TotalOut++
		}
	}
	return &response, nil
}

func (regis *Registration) GetRoom() ([]schema.RoomData, error) {
	return repo.NewRoom().Get()
}
