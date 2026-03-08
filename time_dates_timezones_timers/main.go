package main

import "time"

func ParseDate(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func NewTimerDone(d time.Duration) <-chan time.Time {
	t := time.NewTimer(d)
	return t.C
}
