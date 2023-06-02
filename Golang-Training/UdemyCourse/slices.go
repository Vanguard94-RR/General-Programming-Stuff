package main

import "fmt"

func main() {

	// //
	// keywords := []string{}
	// fmt.Println(keywords)

	// var x []string
	// fmt.Println(x)

	// //initialize slice with value

	// src := []string{"X", "Y", "Z"}
	// fmt.Println(src)

	// // get/set Values by index

	// fmt.Println(src[0])

	// src[1] = "W"

	// fmt.Println(src)

	// // add new

	// src = append(src, "F")
	// fmt.Println(src)

	// //another way to initialize slices with predefined size

	// dest := make([]string, 4)
	// fmt.Println(dest)

	// //copy values from array1 to another
	// copy(dest, src)
	// fmt.Println(dest)

	// src[0] = "A"

	// fmt.Println(src)
	// fmt.Println(dest)

	// s := src[:3]
	// fmt.Println(s)

	gameMap := [][]string{
		make([]string, 5),
		make([]string, 5),
		[]string{"-", "-"},
	}
	fmt.Println(gameMap)

	gameMap[0] = make([]string, 6)
	gameMap[1] = make([]string, 10)

	fmt.Println(gameMap)
	gameMap[1][0] = "W"
	gameMap[1][3] = "X"
	gameMap[1][6] = "Y"
	gameMap[1][9] = "Z"
	fmt.Println(gameMap)
}
