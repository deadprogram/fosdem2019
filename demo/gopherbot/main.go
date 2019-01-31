// Gopherbot is a TinyGo powered furry toy.
//
// This version is designed to be used with a BBC:Microbit microcontroller.
//
// Alongside the following hardware devices:
// - LED antenna
// - Video with APA102 SPI RGB LEDs.
// - 2 buttons to control the visor display
//
// The visor has 4 modes:
// - green
// - red
// - cylon
// - xmas
//
package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/aykevl/tinygo-drivers/apa102"
)

var (
	antenna *machine.GPIO

	visor *apa102.Device
	leds  []color.RGBA

	rg  = false
	pos = 0
	dir = 0
	max = 6
)

const (
	greenVisor = iota
	redVisor
	cylonVisor
	xmasVisor
)

const (
	forward = iota
	backward
)

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 500000,
		Mode:      0})

	ant := machine.GPIO{machine.ADC0}
	ant.Configure(machine.GPIOConfig{Mode: machine.GPIO_OUTPUT})
	antenna = &ant

	left := machine.GPIO{machine.BUTTONA}
	left.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})

	right := machine.GPIO{machine.BUTTONB}
	right.Configure(machine.GPIOConfig{Mode: machine.GPIO_INPUT})

	a := apa102.New(machine.SPI0)
	visor = &a

	leds = make([]color.RGBA, max)

	mode := 0

	go blinker()

	for {
		if !right.Get() {
			mode++
			if mode > 3 {
				mode = 0
			}
		}

		if !left.Get() {
			mode--
			if mode < 0 {
				mode = 3
			}
		}

		switch mode {
		case greenVisor:
			green()
		case redVisor:
			red()
		case cylonVisor:
			cylon()
		case xmasVisor:
			xmas()
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func blinker() {
	for {
		antenna.High()
		time.Sleep(500 * time.Millisecond)
		antenna.Low()
		time.Sleep(500 * time.Millisecond)
	}
}

func red() {
	for i := range leds {
		leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x77}
	}

	visor.WriteColors(leds)
}

func green() {
	for i := range leds {
		leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0x77}
	}

	visor.WriteColors(leds)
}

func xmas() {
	rg = !rg
	for i := range leds {
		rg = !rg
		if rg {
			leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x77}
		} else {
			leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0x77}
		}
	}

	visor.WriteColors(leds)
}

func cylon() {
	if dir == forward {
		pos++
		if pos == max {
			pos = max - 1
			dir = backward
		}
	} else {
		pos--
		if pos < 0 {
			pos = 0
			dir = forward
		}
	}

	for i := range leds {
		if i == pos {
			leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x77}
		} else {
			leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x77}
		}
	}

	visor.WriteColors(leds)
}
