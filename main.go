package main

import (
	"bufio"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func main() {
	for {
		err := runProgram()
		if err != nil {
			log.Printf("Program encountered an error: %v. Restarting...", err)
			time.Sleep(5 * time.Second) // Add a delay before restarting
		}
	}
}

func runProgram() error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	serialPortConfig := &serial.Config{
		Name: config.Port.Name,
		Baud: config.Port.Baud,
	}

	port, err := serial.OpenPort(serialPortConfig)
	if err != nil {
		return err
	}
	defer port.Close()

	reader := bufio.NewReader(port)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
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
