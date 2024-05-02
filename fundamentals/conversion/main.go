package main

import (
	"fmt"
)


// conversion allows a compiler to convert one data type to another data type **
// casting:
// conversion: 



func main() {

	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y,y)     //int
	fmt.Printf("%v of type %T \n", z,z)        //float64

	var m float32 = 43.742

	z = float64(m)
	var a = float32(y)
	fmt.Printf("%v of type %T \n", m, m)  //float32
	fmt.Printf("%v of type %T \n", z,z)     //float64
	fmt.Printf("%v of type %T", a,a)     //float32


	//short declaration operator can be only used in a function
}

