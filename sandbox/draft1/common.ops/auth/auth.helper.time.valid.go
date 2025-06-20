package auth

import "time"

func ValidateExpirationDate(expirationTimeString string) bool {

	// Parse the expiration time string using the default layout
	expirationTime, err := StringToTime(expirationTimeString)
	if err != nil {
		return false // Invalid date format
	}

	return ValidateExpirationDateWithTime(expirationTime)

}

func ValidateExpirationDateWithTime(expirationTime time.Time) bool {

	if expirationTime.IsZero() {
		return false
	}
	// Get the current time
	currentTime := time.Now()

	return expirationTime.After(currentTime)
}
