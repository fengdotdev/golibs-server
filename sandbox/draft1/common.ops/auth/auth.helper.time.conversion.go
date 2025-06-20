package auth

import (
	"fmt"
	"time"
)

func StringToTime(inputTime string) (time.Time, error) {
	if inputTime == "" {
		return time.Time{}, fmt.Errorf("input time string is empty")
	}

	// Parse the input string using the default layout
	parsedTime, err := time.Parse(LayoutDate, inputTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date %v", err)
	}
	return parsedTime, nil
}

func StringToTimeWithLayout(inputTime string, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, inputTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %v", err)
	}
	return parsedTime, nil
}

func TimeToString(inputTime time.Time) string {
	// Format the time to a string using the default layout
	return inputTime.Format(LayoutDate)
}

func TimeToStringWithLayout(inputTime time.Time, layout string) string {
	return inputTime.Format(layout)
}
