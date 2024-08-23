package main

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
	"strings"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Error getting config data: %v", err)
	}

	serialPortConfig := &serial.Config{
		Name: config.Port.Name,
		Baud: config.Port.Baud,
	}

	port, err := serial.OpenPort(serialPortConfig)
	if err != nil {
		log.Fatalf("Error opening serial port: %v", err)
	}

	defer port.Close()

	reader := bufio.NewReader(port)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Could not read line: %v", err)
			continue
		}

		line = strings.TrimSpace(line)

		cmd := config.GetCommand(line)
		if cmd == nil {
			log.Printf("Command not found: %s", line)
			continue
		}

		err = cmd.Run()
		if err != nil {
			log.Printf("Could not run command: %v", err)
		}
	}
}
