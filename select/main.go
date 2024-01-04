package main

import (
	"log"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "function - 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "function - 2s"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	// for {
	// 	log.Printf("%s", <-c1)
	// 	log.Printf("%s", <-c2)
	// }

	for {
		select {
		case msg1 := <-c1:
			log.Printf("%s", msg1)
		case msg2 := <-c2:
			log.Printf("%s", msg2)
		}
	}
}

/*
In the commented infinite for loop, we are accessing the message on first channel which are
coming every 500ms but since the nature of receiving is blocking, we cannot receive the message
on first channel even if the sender is sending to channel continuously. Hence the messages are
beiing printed in the frequency of the message received on the channel having more sleep time
i.e. channel 2.

In the uncommented for loop, we use the select statement to print the message from whatever channel
is ready.

*/
