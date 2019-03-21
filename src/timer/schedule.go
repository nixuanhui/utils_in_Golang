package timer

import (
	"time"
	"sync"
	"fmt"
)

type Schedule interface {
	Run()
	StopRunning()
	AddTask(action func(), timing int64)
}

// default schedule
func NewSchedule() *timeWheel {
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
	running bool

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
	tw.startTick()

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
	tw.running = false

	time.Sleep(10*time.Microsecond)

	tw.stopTick()

	//停 goroutine.
	tw.stopSignal <- struct{}{}
}

func (tw *timeWheel) startTick() {
	if tw.ticker != nil {
		return
	}

	tw.ticker = time.NewTicker(tw.tickDuration)
	fmt.Println("ticking start")

	go func() {
		for {
			select {
			case <-tw.ticker.C:
				tw.pointer = tw.pointer % tw.ticksPerWheel + 1
				if tw.running {
					tw.toDoChan <- tw.GetTasks(tw.pointer)
				}
			case <-tw.stopSignal:
				fmt.Println("schedule stop")
				return
			}
		}
	}()
}

func (tw *timeWheel) stopTick()  {
	if tw.ticker != nil {
		//停 ticker
		tw.ticker.Stop()
		fmt.Println("ticking stop")
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


