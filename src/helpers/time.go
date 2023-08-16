package helpers

import "time"

var DBFormat = "2006-01-02 15:04:05"

func GetTodayStartDate() string {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, currentTime.Location())
	return midnight.Format(DBFormat)
}

func GetThisMonthStartDate() string {
	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, currentTime.Location())
	return firstOfMonth.Format(DBFormat)
}

func GetThisMonthEndDate() string {
	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, currentTime.Location())
	return lastDayOfMonth.Format(DBFormat)
}

func GetTimeNow() string {
	currentTime := time.Now()
	return currentTime.Format(DBFormat)
}
