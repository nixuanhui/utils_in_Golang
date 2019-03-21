package timer

import (
	"testing"
	"time"
	"fmt"
)

func TestSchedule(t *testing.T) {
	a := NewSchedule()
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		fmt.Println("tick...1")
	}, 1)
	a.AddTask(func() {
		fmt.Println("tick...2")
	}, 2)
	a.AddTask(func() {
		fmt.Println("tick...3")
	}, 3)

	a.SetTicksPerWheel(3).SetTickDuration(time.Second)
	go a.Run()
	time.Sleep(6*time.Second)
	a.StopRunning()
	time.Sleep(1*time.Second)
}

func TestScheduleStop(t *testing.T) {
	a := NewSchedule()
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		time.Sleep(2*time.Second)
		fmt.Println("wait 2")
	}, 1)
	a.AddTask(func() {
		fmt.Println("tick...1")
	}, 1)
	a.AddTask(func() {
		fmt.Println("tick...2")
	}, 2)
	a.AddTask(func() {
		fmt.Println("tick...3")
	}, 3)

	a.SetTicksPerWheel(3).SetTickDuration(time.Second)
	go a.Run()
	time.Sleep(3*time.Second)
	a.stopTick()
	time.Sleep(6*time.Second)
	a.StopRunning()
	time.Sleep(1*time.Second)
}



