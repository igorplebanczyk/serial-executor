# Serial Executor

An interface between a serial port and Windows commands. It listens for input from a serial port and executes a command based on the input.

My personal use case is to have an Arduino board with a 4x4 keypad which allows me to change [SignalRGB](https://signalrgb.com/) effects on my PC with a press of a button.

## Features

* **Custom commands**: Define any command to be executed on any input from the serial port
* **Simple configuration**: Define the port name, baud rate, and commands in a single *config.yaml* file
* **Error handling**: If the program encounters a fatal error, it will attempt to restart itself every 5 seconds up to 5 times. This is important for smooth background operation.
* **Silent operation**: Compile with `go build -ldflags="-H windowsgui"` to hide the console window

## Installation and usage

1. Clone this repo
2. Define the port name and baud rate in a *config.yaml* file in the root directory:
    ```yaml
    port:
      name: COM3
      baud: 9600
    ```
3. Define the commands in the same *config.yaml* file:
    ```yaml
      commands:
        - name: "Fire"
          key: "0"
          script: "start signalrgb://effect/apply/Fire?-silentlaunch-"
    
        - name: "Dark Matter"
          key: "1"
          script: "start signalrgb://effect/apply/Dark%20Matter?-silentlaunch-"
    ```
    * `Name` serves no purpose and is just for documentation purposes
    * `Key` is the expected input from the serial port
    * `Script` is the command to be executed
    <br><br>
4. Compile with `go build -ldflags="-H windowsgui"` to hide the console window
    <br><br>
5. Run the executable and optionally add it to the startup folder to have it run on startup

## Notes

* Requires [Go](https://go.dev/doc/install) version `1.22` or higher to be installed
* Windows only
* Since it is designed to run in the background, when the program encounters a fatal error, it will attempt to restart itself every 5 seconds
* Make sure that the device/program that writes to the serial input is ran first, otherwise an `Access Denied` error might appear
* It was designed to solve for a very specific use case of mine, but it is designed to be universal

