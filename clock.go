package main

import "time"

type Clock struct {
	now func() time.Time
}

func NewClock() Clock {
	return NewClockWithNowFunc(time.Now)
}

func NewClockWithNowFunc(now func() time.Time) Clock {
	return Clock{now: now}
}

func (c Clock) NowString() string {
	const layout = "02/01/2006"
	return c.now().Format(layout)
}
