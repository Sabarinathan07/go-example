package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "extra" // This will cause a deadlock since the channel is full	

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages) // This line will not be reached due to the deadlock
}
