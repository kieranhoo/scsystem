package guard

import "time"

func Date(date string) error {
	date = date + " 15:04:05"
	_, err := time.Parse(time.DateTime, date)
	return err
}
