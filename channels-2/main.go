package main

import "log"

func main() {
	c := make(chan string, 2)

	c <- "dog"
	c <- "big dog"
	// c <- "pibull"

	msg := <-c
	log.Printf("%s", msg)

	msg = <-c
	log.Printf("%s", msg)

	// for msg := range c {
	// 	log.Printf("%s", msg)
	// }
}

/*
To demonstrate, the blocking nature of sending and receiving channels

We can write the sender and the receiver in the main go routine. The
process enters into deadlock because the sending of the channel is a
blocking operation and there is no other go routine that is listening
on the channel. Hence, the process enters into a state of deadlock

Error :
	fatal error: all goroutines are asleep - deadlock!

To solve this, we can make buffered channels and the sending operation
will not be blocking the main goroutine until it is completely filled.

It will however become blocking once more messages are sent into it than
its capacity.
*/
