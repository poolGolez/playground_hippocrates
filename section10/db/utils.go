package db

import (
	"fmt"
	"time"
)

func ParseTime(timeString string) time.Time {
	time, err := time.Parse("2006-01-02 15:04:05.000000000+06:00", timeString)
	if err != nil {
		panic(fmt.Sprintf("Error parsing time: %s", timeString))
	}

	return time
}
