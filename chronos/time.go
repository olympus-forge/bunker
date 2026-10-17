package chronos

import (
	"time"
)

type GoTimes struct {
	today    time.Time
	holidays []time.Time
	weekend  map[int]bool
	nexttime time.Time
}

// New returns a new GoTimes instance initialized with the current time and standard weekends.
func New() *GoTimes {
	now := time.Now()
	return &GoTimes{
		today:    now,
		nexttime: now,
		weekend: map[int]bool{
			int(time.Saturday): true,
			int(time.Sunday):   true,
		},
	}
}

// newTime maintains backward compatibility with internal callers.
func newTime() GoTimesImpl {
	return New()
}

func (gt *GoTimes) isWeekend(d time.Time) (w bool) {
	return gt.weekend[int(d.Weekday())]
}

func (gt *GoTimes) SetToday(today *time.Time) *GoTimes {
	if today != nil {
		gt.today = *today
	}
	gt.nexttime = gt.today
	return gt
}

func (gt *GoTimes) SetHolidays(h []time.Time) *GoTimes {
	gt.holidays = append(gt.holidays, h...)
	return gt
}

func (gt *GoTimes) AddWeekDay(sla int) (t time.Time) {
	if sla <= 0 {
		return gt.today
	}
	counted := 0
	for i := 1; counted < sla; i++ {
		target := gt.addDate(0, 0, i)
		if gt.isWeekend(target) {
			continue
		}
		if len(gt.holidays) > 0 && gt.isOffDay(target) {
			continue
		}
		counted++
		gt.nexttime = target
	}
	return gt.nexttime
}

func (gt *GoTimes) isOffDay(d time.Time) bool {
	for _, dayOff := range gt.holidays {
		dayOffLoc := dayOff.In(d.Location())
		if dayOffLoc.Year() == d.Year() && dayOffLoc.Month() == d.Month() && dayOffLoc.Day() == d.Day() {
			return true
		}
	}
	return false
}

func (gt *GoTimes) addDate(year, month, day int) (t time.Time) {
	return gt.today.AddDate(year, month, day)
}

func AddWeekDay(day int, today *time.Time, holidays []time.Time) time.Time {
	return New().SetToday(today).SetHolidays(holidays).AddWeekDay(day)
}
