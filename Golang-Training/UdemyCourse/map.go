// A map is a data type, it maps keys to values
// The zero value of a map is nil
package main

import "fmt"

func main() {
	// init map if you have initial values

	m := map[string]bool{
		"X": true,
		"Y": false,
	}

	fmt.Println("m", m)

	players := make(map[string]int)

	fmt.Println(players)

	// add values to map

	players["Juan"] = 38
	players["Luis"] = 36
	players["Mago"] = 30
	players["Gaby"] = 28

	fmt.Println("Players:", players)

	// Get data from map

	fmt.Println("Players", players["Juan"], players["Mago"])

	// Delete value from map
	delete(players, "Mago")
	fmt.Println("Players:", players["Gaby"])
	// Deleting a key does not affect a map

	// Check if key exist and get value
	// value will be the default value if key does not exist

	v, ok := players["y"]
	fmt.Println("Missing key, v", v, "ok", ok)

	//Map of slices

	mp := map[string][]string{

		"admins": []string{"Luis", "Mago", "Gaby"},
		"users":  []string{"x", "y", "z"},
	}
	fmt.Println("mp", mp)
}
