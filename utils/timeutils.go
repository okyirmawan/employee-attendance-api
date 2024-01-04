package utils

import (
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func ParseDateTimeToString(time time.Time) string {
	return time.Format(dateTimeLayout)
}

func CurrentLocalTime() time.Time {
	localTimeZone, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(localTimeZone)
}

func ConvertTimeToLocalZoneString(inputTime time.Time) (string, error) {
	// Load the target time zone
	targetLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	// Convert the input time to the target time zone
	targetTime := inputTime.In(targetLocation)

	// Format the time as a string
	targetTimeString := targetTime.Format("2006-01-02 15:04:05")

	return targetTimeString, nil
}
