package auth

import "time"

func GetTimeFor90Days() time.Time {
	return GenerateExpirationDateAsTime(90)
}

func GetTimeFor60Days() time.Time {
	return GenerateExpirationDateAsTime(60)
}

func GetTimeFor30Days() time.Time {
	return GenerateExpirationDateAsTime(30)
}

func GetTimeFor7Days() time.Time {
	return GenerateExpirationDateAsTime(7)
}

func GetTimeFor1Day() time.Time {
	return GenerateExpirationDateAsTime(1)
}

func GetTimeFor24Hours() time.Time {
	return GenerateExpirationDateHoursAsTime(24)
}

func GetTimeFor12Hours() time.Time {
	return GenerateExpirationDateHoursAsTime(12)
}

func GetTimeFor1Hour() time.Time {
	return GenerateExpirationDateHoursAsTime(1)
}

func GetTimeFor30Minutes() time.Time {
	return GenerateExpirationDateMinutesAsTime(30)
}

func GetTimeFor15Minutes() time.Time {
	return GenerateExpirationDateMinutesAsTime(15)
}

func GetTimeFor5Minutes() time.Time {
	return GenerateExpirationDateMinutesAsTime(5)
}

func GetTimeFor1Minute() time.Time {
	return GenerateExpirationDateMinutesAsTime(1)
}
