package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:", r)
		}
	}()
	tic := time.NewTicker(0)
	fmt.Println(<-tic.C)
}
