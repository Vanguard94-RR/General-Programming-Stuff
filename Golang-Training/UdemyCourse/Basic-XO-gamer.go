package main

import "fmt"

func main() {

	// initialize game board with empty strings

	xoBoard := [3][3]string{}

	xoBoard = [3][3]string{}

	//variable to carry player name
	player := "X"

	for {
		fmt.Println("Player", player)

		//red row value

		var row int
		fmt.Println("Please enter row number 0, 1 or 2")
		fmt.Scanln(&row)

		if row < 0 || row > 2 {
			fmt.Println("Invalid Number")

			continue
		}
		//red row value

		var column int
		fmt.Println("Please enter column number 0, 1 or 2")
		fmt.Scanln(&column)

		if column < 0 || column > 2 {
			fmt.Println("Invalid Number")

			continue
		}
		// Set value into game board
		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			//check if position already has value
			fmt.Println("Invalid entry: ", row, column, "value", xoBoard[row][column])
			continue
		}
		//Display game board after each move
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		//Check for winning
		for i := 0; i < 3; i++ {
			// check row
			if xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] == player ||
				xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] == player {

				fmt.Println("Winner is player: ", player)
				//
				return
			}
		}
		if xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == player ||
			xoBoard[0][2] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][0] && xoBoard[2][0] == player {
			fmt.Println("Winner is player: ", player)
			//
			return
		}

		//Swap players
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}

	}
}
