package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan string)

	go printString("cat", 500, c)

	// for i := 0; i < 7; i++ {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	log.Printf("%s", msg)
	// 	log.Printf("%d", len(msg))
	// }

	for msg := range c {
		log.Printf("%s", msg)
	}
}

func printString(s string, timeout int, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
	close(c)
}

/*

sending and receiving of message through channels is a blocking
operation

1.	even if it is an infinite loop inside the function and we are
	retriving the message at the end of the main go routine, the
	whole process would end as soon as the message is received
	because the main goroutine ends. This is also true if the times
	the message is received is equal to or less than the number of
	iterations in the loop.

2.	if try getting message from the channel once all the go routines
	have ended, we will see the message:

	fatal error: all goroutines are asleep - deadlock!

	this is because there is no routine that can serve the channel with
	message and the system with permanently wait for it, hence entering
	a deadlock.

	Go detects this problem at runtime rather than compile time.

3. to solve the problem in 2, we can close the channel
	[NOTE: only the sender entity should close the channel and not the
	receiving entity as it has no idea when the messages are going to stop]

	once the sender has closed the connection, and we still receive it using
	the channel, the channel assigns the message default value for its data
	type(here it is an empty string)

4. if the sender stops channel prematurely, it will cause a panic error.

*/
