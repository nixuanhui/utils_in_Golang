package timer

import (
	"time"
	"sync"
	"fmt"
)

type Schedule interface {
	Run()
	StopRunning()
	AddTask(task func(), timing int64)
	GetTasks(timing int64) []func()
	SetTickDuration(td time.Duration) *timeWheel
	SetTicksPerWheel(count int64) *timeWheel
}

// default schedule
func NewSchedule() Schedule {
	tw := &timeWheel{
		tickDuration: 1*time.Second,
		ticksPerWheel: 60,
		toDoChan:make(chan []func(), 10),
		task: make(map[int64][]func()),
		taskLock: &sync.RWMutex{},
		stopSignal:make(chan struct{}, 1),
	}

	return tw
}

func (tw *timeWheel) SetTickDuration(td time.Duration) *timeWheel {
	if td <= 0 {
		return tw
	}
	tw.tickDuration = td
	return tw
}

func (tw *timeWheel) SetTicksPerWheel(count int64) *timeWheel {
	if count <= 0 {
		return tw
	}
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
	stopSignal    chan struct{}
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
	time.Sleep(10*time.Microsecond)
	tw.stopTick()
}

func (tw *timeWheel) startTick() {
	if tw.ticker != nil {
		return
	}

	tw.ticker = time.NewTicker(tw.tickDuration)
	fmt.Println("ticking start")

	for {
		select {
		case <-tw.ticker.C:
			tw.pointer = tw.pointer % tw.ticksPerWheel + 1
			tw.toDoChan <- tw.GetTasks(tw.pointer)
		case <-tw.stopSignal:
			fmt.Println("ticking stop")
			return
		}
	}
}

func (tw *timeWheel) stopTick()  {
	if tw.ticker != nil {
		//停 ticker
		tw.ticker.Stop()
		//停 goroutine.
		tw.stopSignal <- struct{}{}
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
	return schedule.SetTicksPerWheel(365*24)
}
