package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task func()

func main() {
	tasks := *createTasks(5)

	durationStart := time.Now()
	for _, task := range tasks {
		task()
	}
	durationEnd := time.Now()

	duration := durationEnd.Sub(durationStart)
	fmt.Printf("Total Process Time: %.2f second(s)\n", float64(duration.Milliseconds())/1000.0)
	fmt.Printf("Total Process Time: %.2f second(s)\n", duration.Seconds())
}

func createTasks(count int) *[]Task {
	tasks := make([]Task, count)
	for i := range tasks {
		name := fmt.Sprintf("Task %d", i)
		tasks[i] = createTask(name)
	}
	return &tasks
}

func createTask(name string) func() {
	duration := time.Duration(rand.Intn(4)+1) * time.Second

	return func() {
		fmt.Printf("[%s] Executing task in %7.2f second(s)...\n", name, duration.Seconds())
		time.Sleep(duration)
		fmt.Printf("[%s] Task successfully executed...\n", name)
	}
}
