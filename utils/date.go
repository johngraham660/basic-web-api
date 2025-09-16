package utils

import "time"

func ParseDate(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err) // In a production environment, you'd want to handle this error more gracefully
	}
	return t
}
