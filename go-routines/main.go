package main

import (
	// "fmt"
	"log"
	"sync"
	"time"
)

//there are 2 goroutines here - main and the one we have created
//number of goroutines that you can create has to do with how many cores your cpu has

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//Add increments the counter

	go func() {
		printString("sheep", 500)
		wg.Done()
		//wg Done decrements the wait group counter
	}()

	//Wait waits for wait group counter to turn to 0
	wg.Wait()

	// go printString("sheep", 500)
	// go printString("cat", 500)

	//fmt.Scanln() -> this is a fix to solve not ending main goroutine as
	//this is a blocking function that requires user input
}

func printString(s string, timeout int) {
	for i := 1; true; i++ {
		log.Printf("%s", s)
		time.Sleep(time.Millisecond * time.Duration(timeout))
	}
}

//in case we add both lines 13 14 in go routines because all goroutines end once
//the main go routine ends.
