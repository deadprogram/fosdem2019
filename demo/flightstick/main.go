// TinyGo flight stick using analog joystick and 5 buttons
// Outputs data via serial port in very simple space-delimited format
// End of each line of data has "CR-LF" aka 0x13 0x10
// Each update is sent every 50 ms.
package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	// joystick
	stickX := machine.ADC{machine.ADC0}
	stickX.Configure()

	stickY := machine.ADC{machine.ADC1}
	stickY.Configure()

	// buttons
	b1 := machine.GPIO{machine.BUTTON1}
	b1.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	b2 := machine.GPIO{machine.BUTTON2}
	b2.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	b3 := machine.GPIO{machine.BUTTON3}
	b3.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	b4 := machine.GPIO{machine.BUTTON4}
	b4.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	bJoy := machine.GPIO{machine.ADC2}
	bJoy.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT_PULLUP})

	var (
		xPos    uint16
		yPos    uint16
		b1push  bool
		b2push  bool
		b3push  bool
		b4push  bool
		joypush bool
	)

	for {
		xPos = stickX.Get()
		yPos = stickY.Get()
		b1push = !b1.Get()
		b2push = !b2.Get()
		b3push = !b3.Get()
		b4push = !b4.Get()
		joypush = !bJoy.Get()

		println(xPos, yPos, b1push, b2push, b3push, b4push, joypush)

		time.Sleep(time.Millisecond * 50)
	}
}
