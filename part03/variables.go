// Try different type of variable

package main

import "fmt"

func main() {
	// Variables can be declared with an initial value
	var a = "init"
	fmt.Println(a)

	// Multiple variables can be declared at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Variables declared without an initial value are zero-valued
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable in one line
	f := "pizza"
	fmt.Println(f)

}
