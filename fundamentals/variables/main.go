package main

import "fmt"


// we short declaration operator := 


func main() {
	a := 42
	fmt.Println(a)

	b,c,d, e, f := 0, 1,2,3,"happiness"
	fmt.Println(b,c,d,e, f)
}