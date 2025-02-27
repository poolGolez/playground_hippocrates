package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task func()

func main() {
	tasksRef, channelsRef := createTasks(5)
	tasks, channels := *tasksRef, *channelsRef

	durationStart := time.Now()
	for _, task := range tasks {
		go task()
	}

	for _, channel := range channels {
		<-channel
	}

	durationEnd := time.Now()

	duration := durationEnd.Sub(durationStart)
	fmt.Printf("Total Process Time: %.2f second(s)\n", float64(duration.Milliseconds())/1000.0)
	fmt.Printf("Total Process Time: %.2f second(s)\n", duration.Seconds())
}

func createTasks(count int) (*[]Task, *[]chan bool) {
	tasks := make([]Task, count)
	channels := make([]chan bool, count)

	for i := range tasks {
		name := fmt.Sprintf("Task %d", i)
		channel := make(chan bool)

		channels[i] = channel
		tasks[i] = createTask(name, channel)
	}
	return &tasks, &channels
}

func createTask(name string, channel chan bool) Task {
	duration := time.Duration(rand.Intn(5)+1) * time.Second
	return func() {
		fmt.Printf("[%s] Executing task in %7.2f second(s)...\n", name, duration.Seconds())
		time.Sleep(duration)
		channel <- true
		fmt.Printf("[%s] Task successfully executed...\n", name)
	}
}
