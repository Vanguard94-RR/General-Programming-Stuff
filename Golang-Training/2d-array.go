package main

import "fmt"

func main() {

	array1 := [4][4]string{}

	fmt.Println(array1)

	array1 = [4][4]string{
		[4]string{"Juan", "Luis", "Carlos", "Julio"},
		[4]string{"Mago", "Gaby", "Claudia", "Lety"},
		[4]string{"Enero", "Febrero", "Abril", "Diciembre"},
		[4]string{"1992", "1994", "1986", "1984"},
	}
	fmt.Println(array1)
	fmt.Println(array1[0][0], array1[3][3])
}
