package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
}

func main() {
	dones := make([]chan bool, 4)

	for i := 0; i < 4; i++ {
		dones[i] = make(chan bool)
	}

	go greet("1", dones[0])
	go greet("2", dones[1])
	go slowGreet("3", dones[2])
	go greet("I'm done!", dones[3])

	for _, done := range dones {
		<-done
	}
}
