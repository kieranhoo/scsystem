package service

import (
	"scsystem/internal/domain"
	"scsystem/internal/repo"
	"scsystem/pkg/mailers"
)

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

func InsertNewRoom(req *domain.NewRoomRequest) error {
	repositoty := repo.NewRoom()
	return repositoty.Insert(&domain.Room{
		OfficeID:    req.OfficeID,
		Name:        req.Name,
		Description: req.Description,
	})
}

func InsertNewOffice(req *domain.NewOfficeRequest) error {
	repositoty := repo.NewOffice()
	return repositoty.Insert(&domain.Office{
		OrganizationId: req.OrganizationId,
		Name:           req.Name,
		Address:        req.Address,
		Manager:        req.Manager,
		PhoneNumber:    req.PhoneNumber,
	})
}

func InsertNewOrganization(req *domain.NewOrganizationRequest) error {
	repositoty := repo.NewOrganization()
	return repositoty.Insert(&domain.Organization{
		Name:        req.Name,
		Email:       req.Email,
		Head:        req.Head,
		PhoneNumber: req.PhoneNumber,
		Website:     req.Website,
	})
}

func GetAllOrganization() ([]domain.Organization, error) {
	repositoty := repo.NewOrganization()
	return repositoty.Get()
}

func GetAllOffice() ([]domain.OfficeData, error) {
	repositoty := repo.NewOffice()
	return repositoty.Get()
}
