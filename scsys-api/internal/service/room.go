package service

import (
	"scsystem/internal/model"
	"scsystem/internal/repo"
	"scsystem/internal/schema"
	"scsystem/internal/tasks"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"
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
	_user, err := repo.NewUser().GetByID(req.StudentId)
	if err != nil || _user == nil {
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

func (regis *Registration) RegistrationLatest(studentId, roomId string) (*schema.UserRoomData, error) {
	RoomData := new(schema.UserRoomData)
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	if err := conn.Raw(queries.RegistrationLatest, studentId, roomId).Scan(RoomData).Error; err != nil {
		return nil, err
	}
	history, err := repo.NewHistory().Latest(RoomData.Id)
	if err != nil {
		return nil, err
	}
	switch history.ActivityType {
	case "":
		RoomData.ActivityType = "in"
	case "out":
		RoomData.ActivityType = "in"
	case "in":
		RoomData.ActivityType = "out"
	default:
		panic("unknown activity type")
	}
	return RoomData, nil
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

func (regis *Registration) GetHistories(limit string) ([]model.History, error) {
	return repo.NewHistory().GetList(limit)
}

func (regis *Registration) GetHistoriesData(limit string) ([]schema.HistoryData, error) {
	var historyData []schema.HistoryData
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	if err := conn.Raw(queries.HistoryData, limit).Scan(&historyData).Error; err != nil {
		return nil, err
	}
	return historyData, nil
}
