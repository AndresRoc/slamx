// +build dev

//-----------------------------------------------------------------------------
/*

PC Development System PWM Control

*/
//-----------------------------------------------------------------------------

package pwm

import (
	"log"
)

type PWM struct {
	Name string
	Pin  string
	Val  float32
}

// clamp a value from 0.0 to 1.0
func clamp(x float32) float32 {
	if x > 1.0 {
		return 1.0
	}
	if x < 0.0 {
		return 0.0
	}
	return x
}

// Open the PWM channel
func Open(name, pin string, val float32) (*PWM, error) {
	log.Printf("pwm.Open() %s pin=%s\n", name, pin)
	var pwm PWM
	pwm.Name = name
	pwm.Pin = pin
	pwm.Set(val)
	return &pwm, nil
}

// Close the PWM channel
func (pwm *PWM) Close() {
	log.Printf("pwm.Close() %s\n", pwm.Name)
	// TODO ...
}

// Set the PWM value
func (pwm *PWM) Set(val float32) {
	log.Printf("pwm.Set() %s = %f\n", pwm.Name, pwm.Val)
	pwm.Val = clamp(val)
	// TODO ....
}

//-----------------------------------------------------------------------------