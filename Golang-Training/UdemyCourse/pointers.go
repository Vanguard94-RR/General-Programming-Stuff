// Pointer help us send the memory address fo vaariable around
// instead of sending copy from the variable

package main

import "fmt"

func main() {
	name := "Manuel"
	fmt.Println(name)

	// sending a copy of the variable value
	updateCopy(name)
	fmt.Println(name)

	// we send pointer to the variable by adding "&" before it
	UpdateReference(&name)
	fmt.Println(name)

	//Print a pointer will print the memory adress
	fmt.Println("will print the memory address", &name)

	// if parameter is pointer we can send nil as argument
	defaultvaluepointer(nil)

	var Pvalue *string
	Pvalue = &name

	fmt.Println("pointer", *Pvalue)

}

func updateCopy(s string) {
	s = "mark"
}

func UpdateReference(s *string) {
	*s = "mark s"

	n := "name"

	s = &n

	fmt.Println("up fun", &n, "s", s)
}

func defaultvaluepointer(s *string) {
	fmt.Println("nil val", s)
}
