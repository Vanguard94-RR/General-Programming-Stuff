package main

import "fmt"

func main() {
	//define arrays
	// string array
	var arr [4]string

	fmt.Println("array", arr)
	// integer array
	var a [3]int
	fmt.Println("array", a)

	a[0] = 9
	a[1] = 2
	a[2] = 3

	fmt.Println("array", a)

	arr = [4]string{"Juan", "Luis", "Carlos", "Julio"}
	fmt.Println("Reassigned string array", arr)

	var num = [...]int{9, 12, 24, 48, 60}
	fmt.Println("Integer array", num)
	//changing array element value
	num[3] = 19
	fmt.Println("Integer array", num)
	fmt.Println("Integer array", num[2])
}
