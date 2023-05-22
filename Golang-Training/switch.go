package main

import (
	"fmt"
	"time"
)

// func main() {

// 	playername := "Pedro"

// 	switch playername {
// 	case "Juan":
// 		fmt.Println("Name is correct")
// 	case "Manuel":
// 		fmt.Println("Name is not correct")
// 	case "Luis":
// 		fmt.Println("Name is not correct")
// 	case "Julio":
// 		fmt.Println("Name is not correct")
// 	default:
// 		fmt.Println("Terminating")
// 	}
// }
func main() {
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's weekend", time.Now().Weekday())
	default:
		fmt.Println("It's week day", time.Now().Weekday())

	}

}
