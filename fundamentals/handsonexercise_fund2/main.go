package main

import (
	"fmt"
)


// var statement, variable, type ===> gets the default value
const d int = 42          // parallel type system
var c, python, java bool  // static type

var a,b = 1,2.0


//implicit types

func main() {
	var i int

	k := 3

	fmt.Println(i,c,python, java,d )       // 0 false false false 42 
	fmt.Printf("The types are %T, %T \n", a,b)

	fmt.Printf("%T \t %T \t  %T \t %T \t %T \t %T \t \n ", a,b,c,python,java,d)
	fmt.Println(k)
}