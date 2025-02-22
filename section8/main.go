package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Play concurrency with GoRoutines!")
	start := time.Now()

	job1 := createJob("Job 1")
	job2 := createJob("Job 2")

	channel1 := make(chan int64)
	channel2 := make(chan int64)

	go job1(channel1)
	go job2(channel2)

	<-channel1
	<-channel2

	end := time.Now()
	elapsedTime := end.Sub(start)
	fmt.Printf("\n>> Project executed in %.2f seconds\n", elapsedTime.Seconds())
}

func createJob(name string) func(chan int64) {
	return func(channel chan int64) {
		duration := rand.Int63n(5) + 2
		fmt.Printf("Executing %s for %d seconds...\n", name, duration)
		time.Sleep(time.Duration(duration) * time.Second)
		fmt.Printf("Done %s...\n", name)

		channel <- duration
	}
}
