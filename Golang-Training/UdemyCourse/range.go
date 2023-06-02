// Range is used witin the for loop
// iterates over an array, slice or map to get key/value
package main

import "fmt"

func main() {
	// // Range on array
	// arr := [4]string{"Juan", "Luis", "Gaby", "Mago"}

	// for k, v := range arr {

	// 	fmt.Println(k, v)

	// }
	// //Range on slice
	// slc := []string{"Perro", "Gato", "Caballo", "Gallo"}
	// fmt.Println()
	// for key, value := range slc {

	// 	fmt.Println(key, value)
	// }
	//Range on map of slices
	mapOfSlices := map[string][]string{
		"Admins": []string{"Juan", "Luis", "Gaby", "Mago"},
		"Users":  []string{"Perro", "Gato", "Caballo", "Gallo"},
	}
	for key, value := range mapOfSlices {
		fmt.Println("Range", key, value)
		for k, v := range value {
			fmt.Println("Range on Slices", k, v)
		}

	}
}
