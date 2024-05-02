package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

// nested function call from inside out
// helper functions

/*
이렇게도 쓸수있음

x int, y int

x,y int 
*/


func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return  x,y
}

func swap(x,y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func sayHello(){
	fmt.Println("SAY HELLO \n")
}

func main() {
	
	t0 := time.Now()
	fmt.Printf("The time right now is %v \n", t0)
	fmt.Printf("The type of the variable is %T \n", t0)
	fmt.Println(rand.Intn(100))         // [0, n)         inclusive, exclusive
	fmt.Println(math.Pi)
	fmt.Println(math.Log2E, "\n")
	
	fmt.Println(add(4,3))

	sayHello()


	fmt.Println(swap("Hello", "Hi"))

	fmt.Println(split(9))

	// function can take zero or more arguments

}

// any unexported names are not accessible from outside the package  
// this is not public nor private


// every go program is made up of packages
// programs start running in package main
// math/rand

// intn returns, as an int, non negative pseudo random number in the half open 
