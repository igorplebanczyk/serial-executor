package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strings"
	"time"
)

func main() {
	var restartCount int

	for restartCount < 5 {
		err := runProgram()
		if err != nil {
			restartCount++
			log.Printf("Program encountered an error: %v. Restarting...", err)
			time.Sleep(5 * time.Second)
		}
	}
}

func runProgram() error {
	config, err := GetConfig()
	if err != nil {
		return err
	}

	serialPortConfig := &serial.Config{
		Name:        config.Port.Name,
		Baud:        config.Port.Baud,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
		ReadTimeout: time.Second,
	}

	port, err := serial.OpenPort(serialPortConfig)
	if err != nil {
		return fmt.Errorf("error opening serial port: %v", err)
	}
	defer port.Close()

	reader := bufio.NewReader(port)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error reading from serial port: %v", err)
		}

		line = line[:len(line)-1]
		line = strings.TrimSpace(line)

		cmd := config.GetCommand(line)
		if cmd == nil {
			log.Printf("command not found: %s", line)
			continue
		}

		err = cmd.Run()
		if err != nil {
			log.Printf("error running command: %v", err)
		}
	}
}
