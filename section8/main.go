package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job func(chan int64)

func main() {
	fmt.Println("Play concurrency with GoRoutines!")

	jobCount := 20
	start := time.Now()

	jobs := make([]Job, jobCount)
	channels := make([]chan int64, jobCount)
	for i := range jobs {
		name := fmt.Sprintf("Job %d", (i + 1))
		jobs[i] = createJob(name)
		channels[i] = make(chan int64)
	}

	for index, job := range jobs {
		go job(channels[index])
	}

	for _, channel := range channels {
		<-channel
	}

	end := time.Now()
	elapsedTime := end.Sub(start)
	fmt.Printf("\n>> Project executed in %.2f seconds\n", elapsedTime.Seconds())
}

func createJob(name string) func(chan int64) {
	return func(channel chan int64) {
		duration := rand.Int63n(10) + 1
		fmt.Printf(">>> Executing %s [%d seconds]...\n", name, duration)
		time.Sleep(time.Duration(duration) * time.Second)
		fmt.Printf("  <<<Done %s [%d seconds]...\n", name, duration)

		channel <- duration
	}
}
