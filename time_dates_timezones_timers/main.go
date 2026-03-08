package main

import (
	"fmt"
	"time"
)

func ParseDate(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func NewTimerDone(d time.Duration) <-chan time.Time {
	t := time.NewTimer(d)
	return t.C
}

func main() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("load location error:", err)
		return
	}
	t, err := ParseDate("2006-01-02 15:04", "2026-03-08 10:00", loc)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Println("parsed:", t.Format(time.RFC3339))

	<-NewTimerDone(5 * time.Millisecond)
	fmt.Println("timer fired")
}
