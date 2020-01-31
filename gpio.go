package pucker

import (
	"fmt"
	"time"

	"github.com/TLMcNulty/gpio"
	"github.com/stianeikeland/go-rpio"
)

func (b *Pucker) initGPIO() {
	// Set the button pin from pucker.go to pull down.
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		b.GPIOEnabled = false
		return
	} else {
		b.GPIOEnabled = true
	}

	pin := rpio.Pin(ButtonPin)
	pin.PullDown()

	rpio.Close()

	// Button polling
	// Using a pull down because of my hardware
	b.Button = gpio.NewInput(ButtonPin)
	go func() {
		for {
			currentState, err := b.Button.Read()

			if currentState != b.ButtonState && err == nil {
				b.ButtonState = currentState

				if b.Stream != nil {
					if b.ButtonState == 1 {
						fmt.Printf("Transmit start...\n")
						b.TransmitStart()
					} else {
						fmt.Printf("Transmit stop...\n")
						b.TransmitStop()
					}
				}

			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	// then we can do our gpio stuff
	b.OnlineLED = gpio.NewOutput(OnlineLEDPin, false)
	b.ParticipantsLED = gpio.NewOutput(ParticipantsLEDPin, false)
	b.TransmitLED = gpio.NewOutput(TransmitLEDPin, false)
}

func (b *Pucker) LEDOn(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.High()
}

func (b *Pucker) LEDOff(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.Low()
}

func (b *Pucker) LEDOffAll() {
	if b.GPIOEnabled == false {
		return
	}

	b.LEDOff(b.OnlineLED)
	b.LEDOff(b.ParticipantsLED)
	b.LEDOff(b.TransmitLED)
}
