package chronos

import "time"

type GoTimesImpl interface {
	SetToday(today *time.Time) *GoTimes
	SetHolidays(h []time.Time) *GoTimes
	AddWeekDay(sla int) (t time.Time)
}
