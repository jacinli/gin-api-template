package demo

import (
	"fmt"
)

func DemoChannel() {
	channel := make(chan string)

	go func() {
		channel <- "Hello, World!"
	}()

	msg := <-channel
	fmt.Println(msg)
}
