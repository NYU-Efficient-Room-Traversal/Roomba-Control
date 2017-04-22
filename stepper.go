//
// Laser Range Finder
// stepper.go
//
// Cole Smith - css@nyu.edu
// Eric Lin   - eric.lin@nyu.edu
// LICENSE: Apache 2.0
//

package roombaControl

import (
	"fmt"
	"time"
	"github.com/stianeikeland/go-rpio"
)

type Stepper struct {
	pin_step Pin
	delayMillisecond int
	pin_microstep1 Pin
	pin_microstep2 Pin
	pin_microstep3 Pin
}

func (s Stepper) step() {
	s.pin_step.Output()
	s.pin_step.High()
	time.Sleep(delayMillisecond * time.Millisecond)
	s.pin_step.Low()
	time.Sleep(delayMillisecond * time.Millisecond)
}

func (s Stepper) set_full_step() {
	s.pin_microstep1.Output()
	s.pin_microstep1.Low()
	s.pin_microstep2.Output()
	s.pin_microstep2.Low()
	s.pin_microstep3.Output()
	s.pin_microstep3.Low()
}

func (s Stepper) set_half_step() {
	s.pin_microstep1.Output()
	s.pin_microstep1.High()
	s.pin_microstep2.Output()
	s.pin_microstep2.Low()
	s.pin_microstep3.Output()
	s.pin_microstep3.Low()
}

func (s Stepper) set_quater_step() {
	s.pin_microstep1.Output()
	s.pin_microstep1.Low()
	s.pin_microstep2.Output()
	s.pin_microstep2.High()
	s.pin_microstep3.Output()
	s.pin_microstep3.Low()
}

func (s Stepper) set_eighth_step() {
	s.pin_microstep1.Output()
	s.pin_microstep1.High()
	s.pin_microstep2.Output()
	s.pin_microstep2.High()
	s.pin_microstep3.Output()
	s.pin_microstep3.Low()
}

func (s Stepper) set_sixteenth_step() {
	s.pin_microstep1.Output()
	s.pin_microstep1.High()
	s.pin_microstep2.Output()
	s.pin_microstep2.High()
	s.pin_microstep3.Output()
	s.pin_microstep3.High()
}

func initializeStepper(pin_step int, delayMillisecond int, pin_microstep1 int, pin_microstep2 int, pin_microstep3 int) *Stepper {
	return Stepper{rpio.Pin(pin_step), delayMillisecond, rpio.Pin(pin_microstep1), pio.Pin(pin_microstep2), pio.Pin(pin_microstep3)}
}
