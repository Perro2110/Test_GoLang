// Try different type of values

package main

import "fmt"

func main() {
	// String values are surrounded by double quotes
	fmt.Println("go" + "lnag")

	// Integer and floating point numbers are added, subtracted, multiplied, and divided as expected
	fmt.Println("1 + 1 = ", 1+1)
	// Integer division discards the fractional part
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Boolean values are used with boolean operators as expected
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
