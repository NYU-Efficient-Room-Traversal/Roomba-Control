//
// Laser Range Finder
// roomba.go
//
// Cole Smith - css@nyu.edu
// Eric Lin   - eric.lin@nyu.edu
// LICENSE: Apache 2.0
//

package roombaControl

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strconv"
	"time"
)

const (
	SERIAL_INTERFACE = "ttyAMA0"
	BAUD             = 115200
	SPEED            = 200
)

var (
	ser *serial.Port
)

func init() {
	// Open Connection to Serial Port
	c := &serial.Config{Name: SERIAL_INTERFACE, Baud: BAUD}
	var err error
	ser, err = serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Motor Priming
	ModeStart()
	ModeSafe()
	Stop()

	fmt.Println("Roomba is ready")
}

//
// Mode Set Functions
//

func ModeStart() {
	fmt.Println("Starting...")
	write(128)
}

func ModeSafe() {
	fmt.Println("In safe mode")
	write(131)
}

func ModeDriveDirect() {
	fmt.Println("In drive direct mode")
	write(145)
}

//
// Utility Functions
//

func toHex(val int) (int, int) {
	eqBitVal := 0
	if val >= 0 {
		eqBitVal = val
	} else {
		eqBitVal = (1 << 16) + val
	}
	return (eqBitVal >> 8) & 0xFF, eqBitVal & 0xFF
}

func toBytes(val int) []byte {
	return []byte(strconv.Itoa(val))
}

func write(val int) {
	_, err := ser.Write(toBytes(val))
	if err != nil {
		fmt.Println("Write Error: %v", err)
	}
	time.Sleep(25 * time.Millisecond)
}

func read() []byte {
	buf := make([]byte, 128)
	n, err := ser.Read(buf)
	if err != nil {
		fmt.Println("Read Error: %v", err)
	}
	return buf[n:]
}

//
// Sensor Functions
//

// TODO: Implement if needed
func GetStasis() {}

func GetBumps() {}

//
// Drive Functions
//

func drive(velocity, angle int) {
	velHigh, velLow := toHex(velocity)
	radHigh, radLow := toHex(angle)
	write(137)
	write(velHigh)
	write(velLow)
	write(radHigh)
	write(radLow)
}

func Forward() {
	fmt.Println("Forward...")
	drive(SPEED, 0)
}

func Backward() {
	fmt.Println("Backward...")
	drive(SPEED*-1, 0)
}

func Stop() {
	fmt.Println("Stopping...")
	velHigh, velLow := toHex(0)
	radHigh, radLow := toHex(0)
	write(137)
	write(velHigh)
	write(velLow)
	write(radHigh)
	write(radLow)
}

func Turn() {
	fmt.Println("Turning...")
	drive(SPEED, -1)
	// 0.54 is ~ 90 degrees
	time.Sleep(54 * time.Millisecond)
}

func TurnLeft() {
	fmt.Println("Turning...")
	drive(SPEED, 1)
	// 0.54 is ~ 90 degrees
	time.Sleep(54 * time.Millisecond)
}
