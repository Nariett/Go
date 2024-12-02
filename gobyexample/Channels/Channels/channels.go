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

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func pingPong(ping, pong chan string) {
	for {
		select {
		case msg := <-ping:
			fmt.Println("Received:", msg)
			pong <- "pong"
			time.Sleep(2 * time.Second)
		case msg := <-pong:
			fmt.Println("Received:", msg)
			ping <- "ping"
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	//сhannels
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
	fmt.Println()

	//Channel Buffering
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println()

	//Channel Synchronization
	done := make(chan bool, 1)
	go worker(done)

	<-done
	fmt.Println()

	//Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	//Select
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	time.Sleep(5 * time.Second)

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
	fmt.Println("Task")
	//task
	ping := make(chan string)
	pong := make(chan string)

	go pingPong(ping, pong)
	go pingPong(ping, pong)
	ping <- "ping"
	time.Sleep(5 * time.Second)

}
