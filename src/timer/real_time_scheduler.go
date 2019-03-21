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


func (rtcs realTimeCalendarSchedule) Run() {
	rtcs.addtask(rtcs.ticksPerWheel-1, func() {
		//sleep 1秒，以防止程序还未进入到下一年，导致time.Now()不符合预期
		time.Sleep(1 * time.Second)
		setTicksPerWheel(rtcs.timeWheel)
	})

	rtcs.timeWheel.Run()
}

func setTicksPerWheel(schedule *timeWheel)  {
	if time2.IsGapYear(time.Now().Year()) {
		schedule.SetTicksPerWheel(366 * 24 * 60 * 60)
	} else {
		schedule.SetTicksPerWheel(365 * 24 * 60 * 60)
	}
}