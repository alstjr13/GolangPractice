package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary %b \n", adams)
	fmt.Printf("42 as hexadecimal %x \n", adams)

	a,b,c,d,e,f := 0,1,2,3,4,5

	fmt.Printf("%b \n %b \n %b \n %b \n %b \n %b \n", a,b,c,d,e,f)
	fmt.Printf("%v \t %b \t %#X", a,a,a)
	fmt.Printf("%v \t %b \t %#X", b,b,b)
	fmt.Printf("%v \t %b \t %#X", c,c,c)
	fmt.Printf("%v \t %b \t %#X", d,d,d)
	fmt.Printf("%v \t %b \t %#X", e,e,e)
	fmt.Printf("%v \t %b \t %#X", f,f,f)
}