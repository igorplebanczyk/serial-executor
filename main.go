package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	config := &serial.Config{
		Name: "COM3",
		Baud: 115200, // I hope that higher baud rate will reduce input lag
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	defer func(port *serial.Port) {
		err := port.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(port)
}
