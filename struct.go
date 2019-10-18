
package main

import ("fmt")

const u16bitmax = 65535
const kmpmultiple = 1.609

type car struct {
	gas_pedal uint16    
	brake_pedal uint16   
	steering_wheel int16 
	top_speed_kmh float64 
}

// Method from struct
func (c car) kmp() float64 {

	return float64(c.gas_pedal)*(c.top_speed_kmh/u16bitmax)
}
 // Pointer Receiver efficient to used when struct is big 
func (c *car) ptrkmp() float64 {
	c.top_speed_kmh=500

	return float64(c.gas_pedal)*(c.top_speed_kmh/u16bitmax)
}


func main() {
	/*a_car := car{gas_pedal: 16535, 
				 brake_pedal: 0, 
				 steering_wheel: 12562, 
				 top_speed_kmh: 225.0
				}
	*/
	a_car := car{65000,0,12562,225.0}
    fmt.Println(a_car.gas_pedal)
    fmt.Println(a_car.kmp())
    fmt.Println(a_car.ptrkmp())


	//fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
}