package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task func()

func main() {
	tasksRef, doneChannelsRef, errChannelsRef := createTasks(9)
	tasks, doneChannels, errChannels := *tasksRef, *doneChannelsRef, *errChannelsRef

	durationStart := time.Now()
	for _, task := range tasks {
		go task()
	}

	for index := range tasks {
		select {
		case name := <-doneChannels[index]:
			fmt.Printf("[%s] Task successfully executed...\n", name)
		case err := <-(errChannels[index]):
			fmt.Println("An error happened:", err)
		}
	}

	durationEnd := time.Now()

	duration := durationEnd.Sub(durationStart)
	fmt.Printf("Total Process Time: %.2f second(s)\n", duration.Seconds())
}

func createTasks(count int) (*[]Task, *[]chan any, *[]chan error) {
	tasks := make([]Task, count)
	doneChannels := make([]chan any, count)
	errorChannels := make([]chan error, count)

	for i := range tasks {
		name := fmt.Sprintf("Task %d", i+1)
		done := make(chan any)
		err := make(chan error)

		doneChannels[i] = done
		errorChannels[i] = err
		tasks[i] = createTask(name, done, err)
	}
	return &tasks, &doneChannels, &errorChannels
}

func createTask(name string, doneChannel chan any, errChannel chan error) Task {
	duration := time.Duration(rand.Intn(7)+1) * time.Second
	return func() {
		fmt.Printf("[%s] Executing task in %7.2f second(s)...\n", name, duration.Seconds())
		time.Sleep(duration)

		if duration <= time.Duration(4)*time.Second {
			doneChannel <- name
		} else {
			errChannel <- fmt.Errorf("[%s] timed out for %.2f seconds", name, duration.Seconds())
		}
	}
}
