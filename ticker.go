package main

import (
	"fmt"
	"time"
)

func main() {
	tic := time.NewTicker(time.Second * 2)

	done := make(chan bool)
	defer close(done)

	go func() {
		for {
			select {
			case t := <-tic.C:
				fmt.Println("Tick at ", t)
			case <-done:
				fmt.Println("stop ticking")
				tic.Stop()
				return
			}
		}
	}()

	time.Sleep(time.Second * 8)
	done <- true
}
