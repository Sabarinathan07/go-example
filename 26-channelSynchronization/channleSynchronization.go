package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	fmt.Println("starting main...")
	time.Sleep(1 * time.Second) // Simulate some work in main
	fmt.Println("starting worker...")

	go worker(done)
	time.Sleep(1 * time.Second) // Simulate some work in main
	fmt.Println("waiting... in main")

	cc := <-done // Wait for the worker to finish
    fmt.Println(cc, "received from worker")
    fmt.Println("worker finished, main resuming...")

	fmt.Println("worker finished, main exiting")
}
