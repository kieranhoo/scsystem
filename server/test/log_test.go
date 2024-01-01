package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestYMD(t *testing.T) {
	const dateFormat = "2006-01-02 15:04:05"
	dateStr := "2023-01-27 08:09:37"
	timeNow, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	// timeNow := time.Now().UTC()
	YMD := timeNow.Format(time.DateOnly)
	fmt.Print(YMD)
}
