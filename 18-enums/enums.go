package main

import (
	"fmt"
)

type ServerState int

const (
	STARTING ServerState = iota
	RUNNING
	STOPPING
	STOPPED
)

var stateName = map[ServerState]string{
	STARTING: "Starting",
	RUNNING:  "Running",
	STOPPING: "Stopping",
	STOPPED:  "Stopped",
}

func (s ServerState) String() string {
	if name, ok := stateName[s]; ok {
		return name
	}
	return "Unknown"
}

func transition(s ServerState) ServerState {
	switch s {
	case STARTING:
		fmt.Println("Server is starting")
		return RUNNING
	case RUNNING:
		fmt.Println("Server is running")
		return STOPPING
	case STOPPING:
		fmt.Println("Server is stopping")
		return STOPPED
	case STOPPED:
		fmt.Println("Server is stopped")
		return STARTING
	default:
		fmt.Println("Unknown state")
		panic("Invalid state")
	}
}

func main() {
	ns := transition(STARTING)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}
