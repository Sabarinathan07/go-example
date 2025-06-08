package main

import "fmt"

func ping(pings chan<- string, msg string) {
	// Send a message to the pings channel
	fmt.Println("pinging:", msg)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// Receive a message from the pings channel and send it to the pongs channel
	fmt.Println("ponging...", "received", "from pings", pongs)
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	// fmt.Println(<-pongs, "received from pongs")
	//  the above code will not work because the pongs channel is empty at this point
}
