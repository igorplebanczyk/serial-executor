package main

import (
	"bufio"
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

	defer port.Close()

	commands, err := LoadCommands()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(port)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Could not read line:", err)
		}
		line = line[:len(line)-1] // Remove newline character

		cmd := commands.GetCommand(line)
		if cmd == nil {
			log.Println("Command not found:", line)
		}

		err = cmd.Run()
		if err != nil {
			log.Println("Could not run command:", err)
		}
	}
}
