package service

import (
	"scsystem/internal/model"
	"scsystem/internal/repo"
	"scsystem/internal/schema"
	"scsystem/internal/tasks"
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

func (regis *Registration) GetHistoriesData(limit string) ([]schema.HistoryData, error) {
	return repo.NewHistory().GetHistory(limit)
}

func (regis *Registration) GetRoom() ([]schema.RoomData, error) {
	return repo.NewRoom().Get()
}
