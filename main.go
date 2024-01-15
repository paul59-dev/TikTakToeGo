package main

import "fmt"

type Player struct {
	Pseudo string
	Pawn   string
}

func main() {
	var (
		row    int = 3
		column int = 3
	)

	// create a slice of slice == matrix
	matrix := make([][]string, row) // inject number of row into first slice
	for i := range matrix {
		matrix[i] = make([]string, column) // inject number of column into second slice
	}

	// Init the matrix
	// The game board is not formed here
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			matrix[i][j] = "◻️ " // add "◻️" foreach elements of the matrix
		}
	}

	player1 := Player{
		Pseudo: "Paul",
		Pawn:   "🔵",
	}

	player2 := Player{
		Pseudo: "John",
		Pawn:   "🔴",
	}

	fmt.Println(player1)
	fmt.Println(player2)
	// Form the matrix for make the game board
	for i := 0; i < row; i++ { // create row to 3 rows
		for j := 0; j < column; j++ { // create column
			fmt.Printf("%s", matrix[i][j]) // [0][0], [0][1], [0][2] == 3, finish the loop j
		}
		fmt.Println()
	}

}
