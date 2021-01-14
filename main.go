package main

import (
	"log"
	"time"
	"uart-relay/uart"
)

func main() {
	port1, err := uart.NewUart("dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}

	port2, err := uart.NewUart("dev/ttyUSB1")
	if err != nil {
		log.Fatal(err)
	}

	r, err := uart.NewRelay(port1, port2)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 10)

	err = r.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
