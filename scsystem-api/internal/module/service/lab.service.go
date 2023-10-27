package service

import (
	"qrcheckin/internal/config"
	"qrcheckin/internal/module/model"
	"qrcheckin/internal/module/schema"
	"qrcheckin/internal/module/tasks"
	"qrcheckin/internal/module/types"
	"qrcheckin/pkg/database"
	"qrcheckin/pkg/x/worker"
)

type Registration struct {
	repo types.IRegistration
}

func NewLab() types.ILabService {
	return &Registration{
		repo: model.NewRegistration(),
	}
}

func (regis *Registration) RegisterLab(req *schema.RegistrationLabRequest) error {
	_user, err := model.NewUser().GetByID(req.StudentId)
	if err != nil || _user == nil {
		go func() {
			// if err := worker.Execute(
			// 	tasks.DefaultQueue,
			// 	tasks.WorkerSaveUser,
			// 	tasks.SaveUser,
			// 	req.StudentId,
			// 	req.FirstName,
			// 	req.LastName,
			// 	req.PhoneNumber,
			// 	req.Email,
			// ); err != nil {
			// 	panic(err.Error())
			// }
			if err := worker.Exec(config.CriticalQueue, worker.NewTask(
				tasks.WorkerSaveUser,
				types.Users{
					Id:          req.StudentId,
					FirstName:   req.FirstName,
					LastName:    req.LastName,
					PhoneNumber: req.PhoneNumber,
					Email:       req.Email,
				}),
			); err != nil {
				panic(err.Error())
			}
		}()
	}

	// return worker.Execute(
	// 	tasks.DefaultQueue,
	// 	tasks.WorkerSaveRegistration,
	// 	tasks.SaveRegistration,
	// 	req.RegistrationTime,
	// 	req.Supervisor,
	// 	req.StartDay,
	// 	req.EndDay,
	// 	req.StudentId,
	// 	req.RoomId,
	// )

	return worker.Exec(config.CriticalQueue, worker.NewTask(
		tasks.WorkerSaveRegistration,
		types.Registration{
			RegistrationTime: req.RegistrationTime,
			Supervisor:       req.Supervisor,
			StartDay:         req.StartDay,
			EndDay:           req.EndDay,
			UserID:           req.StudentId,
			RoomId:           req.RoomId,
		},
	))
}

func (regis *Registration) RegistrationLatest(studentId, roomId string) (*schema.UserLabData, error) {
	labData := new(schema.UserLabData)
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	if err := conn.Raw(`
	SELECT registration.id AS id,  registration_time, user_id AS student_id, start_day, end_day, first_name, last_name, email, phone_number, org_name, office_name, room_name, supervisor
	FROM registration 
		JOIN users ON users.id = registration.user_id
    	JOIN (
			SELECT organization.name AS org_name, office.name AS office_name, room.name AS room_name, office.id 
			FROM room 
				JOIN office ON room.office_id=office.id 
				JOIN organization ON organization.id=office.organization_id
		) AS rooms ON rooms.id = registration.room_id
	WHERE user_id = ? AND room_id = ?
	ORDER BY registration.id DESC LIMIT 1;`, studentId, roomId).Scan(labData).Error; err != nil {
		return nil, err
	}
	history, err := model.NewHistory().Latest(labData.Id)
	if err != nil {
		return nil, err
	}
	switch history.ActivityType {
	case "":
		labData.ActivityType = "in"
	case "out":
		labData.ActivityType = "in"
	case "in":
		labData.ActivityType = "out"
	default:
		panic("unknown activity type")
	}
	return labData, nil
}

// func GetUser(id string) (*model.Users, error) {
// 	return new(model.Users).GetByID(id)
// }

func (regis *Registration) SaveActivityType(req *schema.CheckInRequest) error {
	// return worker.Execute(
	// 	tasks.DefaultQueue,
	// 	tasks.WorkerSaveActivityType,
	// 	tasks.SaveActivityType,
	// 	req.RegistrationId,
	// 	req.AdminId,
	// 	req.ActivityType,
	// )
	return worker.Exec(config.CriticalQueue, worker.NewTask(
		tasks.WorkerSaveActivityType,
		types.History{
			RegistrationId: req.RegistrationId,
			AdminId:        req.AdminId,
			ActivityType:   req.ActivityType,
		},
	))
}

func (regis *Registration) GetHistories(limit string) ([]types.History, error) {
	return model.NewHistory().GetList(limit)
}

func (regis *Registration) GetHistoriesData(limit string) ([]schema.HistoryData, error) {
	var historyData []schema.HistoryData
	conn, err := database.Connection()
	if err != nil {
		return nil, err
	}
	if err := conn.Raw(`
	SELECT history.id, activity_type, time, admin_id, registration_time, supervisor, 
		user_id, start_day, end_day, first_name, last_name, users.email, users.phone_number, 
		room.name AS room_name, office.name AS office_name, organization.name AS org_name  
	FROM history
		JOIN registration ON registration_id = registration.id
		JOIN users ON users.id = user_id
		JOIN room ON room.id = room_id
		JOIN office ON office.id = office_id
		JOIN organization ON organization.id = organization_id
	ORDER BY history.id DESC LIMIT ?;`, limit).Scan(&historyData).Error; err != nil {
		return nil, err
	}
	return historyData, nil
}
