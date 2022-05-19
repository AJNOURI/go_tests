package golaunch

import (
	"time"
)

func GetDaysSinceGoLaunch() float64 {
	t0 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t1 := time.Now()
	duration := t1.Sub(t0)
	return duration.Hours() / 24
}
