package timer

import (
	"testing"
	"fmt"
)

func TestNewRealTimeCalendarScheduler(t *testing.T) {
	rts := NewRealTimeCalendarScheduler()
	fmt.Println(rts.(*realTimeCalendarSchedule).pointer)
}