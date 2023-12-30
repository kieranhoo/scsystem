package service

import (
	"scsystem/internal/model"
	"scsystem/internal/repo"
	"scsystem/internal/schema"
	"scsystem/internal/tasks"
	"scsystem/pkg/mailers"
	"scsystem/pkg/worker"
)

func HealthCheck() error {
	return worker.Exec(
		tasks.CriticalQueue,
		worker.NewTask(
			tasks.WorkerHealthCheck,
			1,
		),
	)
}

func Email() (string, error) {
	mailers.ConfirmRegistrationRoom("iduchungho@gmail.com", mailers.ConfirmEmail{
		Name:        "Đức Hưng Hồ",
		RoomNumber:  "1",
		Date:        "30/12/2023",
		Time:        "7h-9h",
		StartTime:   "2020-08-05 18:19:00",
		EndTime:     "2020-08-05 18:19:00",
		Email:       "tarzaines@gmail.com",
		FullName:    "ad",
		Position:    "admin",
		Information: "ist Director",
	})
	return "", nil
}

func InsertNewRoom(req *schema.NewRoomRequest) error {
	repositoty := repo.NewRoom()
	return repositoty.Insert(&model.Room{
		OfficeID:    req.OfficeID,
		Name:        req.Name,
		Description: req.Description,
	})
}

func InsertNewOffice(req *schema.NewOfficeRequest) error {
	repositoty := repo.NewOffice()
	return repositoty.Insert(&model.Office{
		OrganizationId: req.OrganizationId,
		Name:           req.Name,
		Address:        req.Address,
		Manager:        req.Manager,
		PhoneNumber:    req.PhoneNumber,
	})
}

func InsertNewOrganization(req *schema.NewOrganizationRequest) error {
	repositoty := repo.NewOrganization()
	return repositoty.Insert(&model.Organization{
		Name:        req.Name,
		Email:       req.Email,
		Head:        req.Head,
		PhoneNumber: req.PhoneNumber,
		Website:     req.Website,
	})
}

func GetAllOrganization() ([]model.Organization, error) {
	repositoty := repo.NewOrganization()
	return repositoty.Get()
}

func GetAllOffice() ([]schema.OfficeData, error) {
	repositoty := repo.NewOffice()
	return repositoty.Get()
}
