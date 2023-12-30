package test

import (
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
