package main

import (
	"fmt"
	"time"
)

func main() {
	// To create a channel, use `make(chan type)`.
    // Channel can only accept values of the specified type:
	message := make(chan string)
	// To send a value to a channel,
    // use the `channel <-` syntax.
    // Let's send "ping":
	go func() {
		fmt.Println("B: Sending message...")
		message <- "ping"
		fmt.Println("B: Message sent!")
		}()
	// To receive a value from a channel,
    // use the `<-channel` syntax.
    // Let's receive "ping" and print it:
	fmt.Println("A: doing something...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Receiving message...")
	<-message                          //  (3)

    fmt.Println("A: Message received!")
    time.Sleep(100 * time.Millisecond)
}