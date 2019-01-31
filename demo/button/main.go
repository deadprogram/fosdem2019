package main

import (
	"machine"
	"time"
)

const buttonPin = 8

func main() {
	led := machine.GPIO{machine.LED}
	led.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})

	button := machine.GPIO{buttonPin}
	button.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})

	for {
		if button.Get() {
			led.Low()
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
