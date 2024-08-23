package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strings"
)

func main() {
	config := &serial.Config{
		Name: "COM3",
		Baud: 115200,
	}

	// Attempt to open the serial port
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Error opening serial port: %v", err)
	}

	defer func() {
		if err := port.Close(); err != nil {
			log.Printf("Error closing serial port: %v", err)
		}
	}()

	// Load commands from YAML file
	commands, err := LoadCommands()
	if err != nil {
		log.Fatalf("Failed to load commands: %v", err)
	}

	reader := bufio.NewReader(port)
	fmt.Println("Listening for commands...")

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Could not read line: %v", err)
			continue
		}

		line = strings.TrimSpace(line) // Remove newline and any extra spaces

		fmt.Printf("Received command: %s\n", line)
		cmd := commands.GetCommand(line)
		if cmd == nil {
			log.Printf("Command not found: %s", line)
			continue
		}

		fmt.Printf("Running command: %s\n", cmd.Name)
		err = cmd.Run()
		if err != nil {
			log.Printf("Could not run command: %v", err)
		}
	}
}
