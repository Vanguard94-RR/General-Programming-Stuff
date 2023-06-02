package main

import (
	"fmt"
)

func main() {
	//loops

	// 	for i := 0; i <= 50; i++ {
	// 		fmt.Print("Hello", i)
	// 	}
	// }

	// 	fmt.Println("for x <=5")

	// 	x := 1

	// 	for x <= 9 {
	// 		fmt.Println("x value before loop", x)
	// 		x++

	// 		fmt.Println("x value after loop", x)
	// }
	// 	for i := 0; i <= 4; i++ {
	// 		fmt.Println()
	// 		fmt.Println("Iteration # ", i)
	// 		for a := 0; a <= 5; a++ {
	// 			if a == 5 {
	// 				break //continue //

	// 			}
	// 			fmt.Print(a, " ")

	// 		}
	// 	}
	// 	fmt.Println()
	// }
breakpoint:
	for x := 0; x < 6; x++ {
		for y := 0; y < 5; y++ {
			if x == y && y == 5 {
				break breakpoint
			}
			fmt.Println("x,y", x, y)
		}
	}
}
