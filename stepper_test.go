//
// Laser Range Finder
// stepper_test.go
//
// Cole Smith - css@nyu.edu
// Eric Lin   - eric.lin@nyu.edu
// LICENSE: Apache 2.0
//

package roombaControl

import "testing"

func TestInitializeStepper(t *testing.T){
	s := InitializeStepper(18, 5, 24, 17, 25)
	
	count := 0
	for count < 1600 {
		count += 1
		s.step()
	}
	
}