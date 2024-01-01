package mailers

import (
	"bytes"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

var (
	_email       string
	_appPassword string
)

func Config(email, appPassword string) {
	_email = email
	_appPassword = appPassword
}

type ConfirmEmail struct {
	Name        string
	RoomNumber  string
	Date        string
	Time        string
	StartTime   string
	EndTime     string
	Email       string
	FullName    string
	Position    string
	Information string
}

type RemiderData struct {
	Name          string
	OldExpiryDate string
	LabName       string
	RenewalLink   string
}

func ConfirmRegistrationRoom(to string, data ConfirmEmail) error {
	t, err := template.ParseFiles("pkg/template/confirm-en.html")
	if err != nil {
		log.Fatal(err)
	}
	var body bytes.Buffer
	err = t.Execute(&body, data)
	if err != nil {
		return err
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", _email)
	mail.SetHeader("To", to)
	// mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mail.SetHeader("Subject", "Confirmation of Lab Room Reservation")
	mail.SetBody("text/html", body.String())
	// mail.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, _email, _appPassword)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}

func RemiderRenewTime(to string, data RemiderData) error {
	t, err := template.ParseFiles("pkg/template/reminder.html")
	if err != nil {
		log.Fatal(err)
	}
	var body bytes.Buffer
	err = t.Execute(&body, data)
	if err != nil {
		return err
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", _email)
	mail.SetHeader("To", to)
	// mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mail.SetHeader("Subject", "Reminder to Renew Laboratory Usage Timen")
	mail.SetBody("text/html", body.String())
	// mail.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, _email, _appPassword)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
