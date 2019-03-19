package timer

import (
	"time"
	"sync"
)

type Schedule interface {
	Run()
	StopRunning()
	startTick()
	stopTick()
	AddTask(task func(), timing int64)
	GetTasks(timing int64) []func()
	SetTickDuration(td time.Duration) *timeWheel
	SetTicksPerWheel(count int64) *timeWheel
}

// default
func NewSchedule() Schedule {
	tw := &timeWheel{
		tickDuration: 1*time.Second,
		ticksPerWheel: 60,
		toDoChan:make(chan []func(), 10),
		task: make(map[int64][]func()),
		taskLock: &sync.RWMutex{},
	}

	return tw
}

func (tw *timeWheel) SetTickDuration(td time.Duration) *timeWheel {
	tw.tickDuration = td
	return tw
}

func (tw *timeWheel) SetTicksPerWheel(count int64) *timeWheel {
	tw.ticksPerWheel = count
	return tw
}

type timeWheel struct {
	ticker        *time.Ticker
	tickDuration  time.Duration
	ticksPerWheel int64
	pointer       int64
	task          map[int64][]func()
	taskLock      *sync.RWMutex
	toDoChan      chan []func()
}

func (tw *timeWheel) Run()  {
	go tw.startTick()
	var task []func()
	for{
		select {
		case task = <- tw.toDoChan:
			for _, fc := range task {
				go func(f func()) {
					f()
				}(fc)
			}
		}
	}
}

func (tw *timeWheel) StopRunning() {
	time.Sleep(100*time.Microsecond)
	tw.stopTick()
}

func (tw *timeWheel) startTick() {
	if tw.ticker != nil {
		return
	}

	tw.ticker = time.NewTicker(tw.tickDuration)

	for {
		select {
		case <-tw.ticker.C:
			tw.pointer = tw.pointer % tw.ticksPerWheel + 1
			tw.toDoChan <- tw.GetTasks(tw.pointer)
		}
	}
}

func (tw *timeWheel) stopTick()  {
	if tw.ticker != nil {
		tw.ticker.Stop()
	}
}

func (tw *timeWheel) AddTask(action func(), timing int64) {
	tw.addtask(timing, action)
}

func (tw *timeWheel) addtask(key int64, action func()) {
	tw.taskLock.Lock()
	defer tw.taskLock.Unlock()
	tw.task[key] = append(tw.task[key], action)
}

func (tw *timeWheel) GetTasks(timing int64) []func() {
	tw.taskLock.RLock()
	defer tw.taskLock.RUnlock()
	return tw.task[timing]
}

func NewRealTimeCalendar() Schedule {
	schedule := NewSchedule()
	return schedule.SetTicksPerWheel(365)
}
