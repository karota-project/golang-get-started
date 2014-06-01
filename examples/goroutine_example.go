package main

import (
	"fmt"
	"time"
)

func main() {
	done := false
	/* another goroutine*/
	go func() {
		time.Sleep(time.Second * 3)
		done = true
	}()

	/* main goroutine*/
	for !done {
		time.Sleep(time.Millisecond)
	}

	fmt.Println("done!")

}
