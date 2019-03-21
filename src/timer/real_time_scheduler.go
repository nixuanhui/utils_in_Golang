package timer

import (
	"time"
	"time2"
)

type realTimeCalendarSchedule struct {
	*timeWheel
}

func NewRealTimeCalendarScheduler() Schedule {
	schedule := &realTimeCalendarSchedule{NewSchedule()}

	schedule.setTicksPerWheel()
	schedule.setPointer()
	schedule.startTick()

	return schedule
}


func (rtcs realTimeCalendarSchedule) Run() {
	rtcs.addtask(rtcs.ticksPerWheel-1, func() {
		//sleep 1秒，以防止程序还未进入到下一年，导致time.Now()不符合预期
		time.Sleep(1 * time.Second)
		rtcs.setTicksPerWheel()
	})

	rtcs.timeWheel.Run()
}

func (rtcs *realTimeCalendarSchedule) setTicksPerWheel()  {
	if time2.IsGapYear(time.Now().Year()) {
		rtcs.SetTicksPerWheel(366 * 24 * 60 * 60)
	} else {
		rtcs.SetTicksPerWheel(365 * 24 * 60 * 60)
	}
}

func (rtcs *realTimeCalendarSchedule) setPointer() {
	rtcs.pointer = int64(time.Now().YearDay() * 24 * 60 + time.Now().Hour() * 60 * 60 +
		time.Now().Minute() * 60 + time.Now().Second())
}