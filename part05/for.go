package main

import "fmt"

func main() {

	// while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// for range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// infinite loop || while true
	for {
		fmt.Println("loop")
		break
	}

	// study case
	for n := range 6 {
		if n%2 == 0 { // 	if n is even,
			continue //  skip the rest of the loop body
		} //	and continue with the next iteration
		fmt.Println(n)
	}
}
