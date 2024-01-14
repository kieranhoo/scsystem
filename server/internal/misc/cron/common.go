package cron

import (
	"log"
	"scsystem/internal/repo"
	"scsystem/pkg/mailers"
)

func ScanInValidDate() {
	//TODO: get data before 1 week
	regisRepo := repo.NewRegistration()
	registration, err := regisRepo.OneWeekBefore()
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
					regisRepo.DeleteById(r.Id)
					log.Fatal(r.Id)
				}
			}
		}
	}
}
