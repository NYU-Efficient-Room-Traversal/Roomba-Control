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
	SERIAL_INTERFACE = "/dev/ttyAMA0"
	BAUD             = 115200
)

var (
	ser *serial.Port
)

func init() {
	c := &serial.Config{Name: SERIAL_INTERFACE, Baud: BAUD}
	var err error
	ser, err = serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
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
	fmt.Println("In safe mode")
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

func Forward() {}

func Backward() {}

func Stop() {}

func Turn() {}
