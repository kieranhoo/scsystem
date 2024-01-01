package test

import (
	"log"
	"scsystem/internal/repo"
	"scsystem/pkg/mailers"
	"testing"
)

func TestConfirmEmail(t *testing.T) {
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
}

func TestScanInValidDate(t *testing.T) {
	//TODO: get data before 1 week
	registration, err := repo.NewRegistration().OneWeekBefore()
	//TODO: send email if date invalid
	if err == nil {
		historyRepo := repo.NewHistory()
		userRepo := repo.NewUser()
		roomRepo := repo.NewRoom()
		for _, r := range registration {
			if r.Id != "" {
				history, err := historyRepo.Latest(r.Id)
				if err == nil && history.ActivityType == "" {
					// send email
					user, _ := userRepo.GetByID(r.UserID)
					room, _ := roomRepo.GetByID(r.RoomId)
					mailers.RemiderRenewTime(user.Email, mailers.RemiderData{
						Name:          user.FirstName + " " + user.LastName,
						OldExpiryDate: r.EndDay,
						LabName:       room.Name,
						RenewalLink:   "https://forms.gle/4HW9QtrSeJyN99vB6",
					})
					log.Fatal("done sending email")
				}
			}
		}
	}
}
