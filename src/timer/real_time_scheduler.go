package timer

import (
	"time"
	"time2"
)

type realTimeCalendarSchedule struct {
	*timeWheel
}

func NewRealTimeCalendarScheduler() Schedule {
	schedule := NewSchedule()

	setTicksPerWheel(schedule)

	schedule.pointer = int64(time.Now().YearDay() * 24 * 60 + time.Now().Hour() * 60 * 60 + time.Now().Minute() * 60 + time.Now().Second())
	schedule.startTick()

	return &realTimeCalendarSchedule{schedule}
}


func (rtcs *realTimeCalendarSchedule) Run() {
	rtcs.addtask(rtcs.ticksPerWheel-1, func() {
		setTicksPerWheel(rtcs.timeWheel)
	})
}

func setTicksPerWheel(schedule *timeWheel)  {
	if time2.IsGapYear(time.Now().Year()) {
		schedule.SetTicksPerWheel(366 * 24 * 60 * 60)
	} else {
		schedule.SetTicksPerWheel(365 * 24 * 60 * 60)
	}
}