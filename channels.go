package main

import (
	"fmt"
	"time"
)

var pongCh = make(chan string)
var pingCh = make(chan string)

func ping() {
	for /* ever */ {
		fmt.Println(<-pingCh)
		pongCh <- "ping"
	}
}

func pong() {
	for /* ever */ {
		fmt.Println(<-pongCh)
		pingCh <- "pong"
	}
}

func main() {

	go ping()
	go pong()

	pingCh <- "pong"

	time.Sleep(time.Nanosecond)
}
