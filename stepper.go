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
	"time"
	"github.com/stianeikeland/go-rpio"
)

type Stepper struct {
	pin_step rpio.Pin
	delayMillisecond float64
	pin_direction rpio.Pin
	pin_microstep1 rpio.Pin
	pin_microstep2 rpio.Pin
	pin_microstep3 rpio.Pin
}

func (s Stepper) Spin() {
	count := 0
	for count < 1600 {
		s.Step()
		count += 1
	}
}

func (s Stepper) Step() {
	s.pin_step.Output()
	s.pin_step.High()
	time.Sleep(time.Duration(s.delayMillisecond) * time.Millisecond)
	s.pin_step.Low()
	time.Sleep(time.Duration(s.delayMillisecond) * time.Millisecond)
}

/*

Here's the exported State constant from rpio
// State of pin, High / Low
const (
	Low State = iota
	High
)

*/

func (s Stepper) set_direction(state rpio.State) {
	s.pin_direction.Output()
	s.pin_direction.Write(state)
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

func InitializeStepper(pin_step int, delayMillisecond float64, pin_direction, pin_microstep1 int, pin_microstep2 int, pin_microstep3 int) Stepper {
	return Stepper{rpio.Pin(pin_step), delayMillisecond, rpio.Pin(pin_direction), rpio.Pin(pin_microstep1), rpio.Pin(pin_microstep2), rpio.Pin(pin_microstep3)}
}
