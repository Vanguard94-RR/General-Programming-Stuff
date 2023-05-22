package main

import "fmt"

func main() {

	//define string variables
	var name string //Define variable type
	name = "Hello Juan"
	fmt.Println("Variable content", name)
	// reassing variable
	name = "Hello Manuel"
	fmt.Println("Variable content", name)
	name = "Hello Cortes"
	fmt.Println("Variable content", name)
	//Another way to define variables
	name1 := "var new name"
	fmt.Println(name1)

	score := 10
	fmt.Println(score)

	age1 := 38
	name3 := "Juan Manuel Cortes"
	fmt.Println(name3, "has", age1, "years old")

	var name4 string
	var age2 int
	var floatNum float32
	var boolean bool
	fmt.Println("String default value", name4)
	fmt.Println("Integer default value", age2)
	fmt.Println("Float default value", floatNum)
	fmt.Println("Boolean default value", boolean)

	//Define constant

	const gameName int = 1984

	fmt.Println("Star Wars", gameName)

	// gameName = "X-wing"
	// fmt.Println("Star Wars", gameName)

}
