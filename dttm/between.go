package dttm

import (
	"time"
)

//BetweenTwoDates is a function that checks if a given date is between a start date and an end date.
func BetweenTwoDates(target, start, end time.Time) bool {
	return (start.Before(target) || start.Equal(target)) && end.After(target)
}
