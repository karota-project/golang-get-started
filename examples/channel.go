package main

import (
	"fmt"
	"time"
)

func rcvCannel(_ch chan bool) {
	/* reciever */
	for done := range _ch {
		fmt.Println(time.Now(), " : ", done)
	}
}
func main() {
	/* make channel*/
	ch := make(chan bool)

	go rcvCannel(ch)

	done := false

	for i := 0; !done; i++ {
		if i == 5 {
			done = true
		}
		/* sender*/
		ch <- done

		time.Sleep(time.Second * 1)
	}

	fmt.Println("done!")
}
