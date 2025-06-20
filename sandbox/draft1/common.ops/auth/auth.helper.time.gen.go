package auth

import "time"

func GenerateExpirationDate(days int) string {
	experationTime := GenerateExpirationDateAsTime(days)
	return TimeToString(experationTime)
}

func GenerateExpirationDateAsTime(days int) time.Time {
	currentTime := time.Now()
	expirationTime := currentTime.AddDate(0, 0, days)
	return expirationTime
}

func GenerateExpirationDateHours(hours int) string {
	expirationTime := GenerateExpirationDateHoursAsTime(hours)
	return TimeToString(expirationTime)
}

func GenerateExpirationDateHoursAsTime(hours int) time.Time {
	currentTime := time.Now()
	expirationTime := currentTime.Add(time.Duration(hours) * time.Hour)
	return expirationTime
}

func GenerateExpirationDateMinutes(minutes int) string {
	expirationTime := GenerateExpirationDateMinutesAsTime(minutes)
	return TimeToString(expirationTime)
}

func GenerateExpirationDateMinutesAsTime(minutes int) time.Time {
	currentTime := time.Now()
	expirationTime := currentTime.Add(time.Duration(minutes) * time.Minute)
	return expirationTime
}
