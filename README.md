# Serial Input - CMD Interface

Allows for automatic execution of Windows CMD commands based on serial input.

My personal use case is to have an Arduino board with a 4x4 keypad which allows me to change [SignalRGB](https://signalrgb.com/) effects on my PC with the press of a button.

## Installation and usage

1. Clone this repo
    <br><br>
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

### Important notes

* Requires [Go](https://go.dev/doc/install) version 1.22+ to be installed
* Designed specifically for Windows, might work on other operating systems with slight modifications
* Since it is designed to run in the background, when the program encounters a fatal error, it will attempt to restart itself every 5 seconds
* Make sure that the device/program that writes to the serial input is ran first, otherwise "Access Denied" error might appear**

## TODO
* Add better logging and error handling
* Add flags to switch the behavior of the program when encountering an error

