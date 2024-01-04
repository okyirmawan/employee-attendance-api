package utils

import (
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func ParseStringToDateTime(dateTimeString string) time.Time {
	parsedTime, err := time.Parse(dateTimeLayout, dateTimeString)

	if err != nil {
		return time.Now()
	}

	return parsedTime
}

func ParseDateTimeToString(time time.Time) string {
	return time.Format(dateTimeLayout)
}

func CurrentLocalTime() time.Time {
	localTimeZone, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(localTimeZone)
}
