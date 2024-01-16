package main

import (
	"fmt"
	"paul/tictactoe/packages/player"
)

func main() {
	var (
		row        int  = 3
		column     int  = 3
		playerTour bool = true // true => player 1, false => player 2
		mtxX, mtxY int
	)

	// create a slice of slice == matrix
	matrix := make([][]string, row) // inject number of row into first slice
	for i := range matrix {
		matrix[i] = make([]string, column) // inject number of column into second slice
	}

	game := player.GetPlayer()
	firstPlayer := game.Players[0]
	secondPlayer := game.Players[1]
	fmt.Println(firstPlayer.Pseudo)

	// Init the matrix
	// The game board is not formed here
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			matrix[i][j] = "◻️ " // add "◻️" foreach elements of the matrix
		}
	}

	// Start the game
	for {
		// Form the matrix for make the game board
		for i := 0; i < row; i++ { // create row to 3 rows
			for j := 0; j < column; j++ { // create column
				fmt.Printf("%s", matrix[i][j]) // [0][0], [0][1], [0][2] == 3, finish the loop j
			}
			fmt.Println()
		}

		if playerTour {
			fmt.Println("Tour de ", firstPlayer.Pseudo)
			fmt.Println("Entrer: ")
			fmt.Print("Entrer nombre X: ")
			fmt.Scan(&mtxX)
			fmt.Print("Entrer nombre Y: ")
			fmt.Scan(&mtxY)

			matrix[mtxX][mtxY] = firstPlayer.Pawn
			playerTour = false
		} else {
			fmt.Println("Tour de ", secondPlayer.Pseudo)
			fmt.Println("Entrer: ")
			fmt.Print("Entrer nombre X: ")
			fmt.Scan(&mtxX)
			fmt.Print("Entrer nombre Y: ")
			fmt.Scan(&mtxY)

			matrix[mtxX][mtxY] = secondPlayer.Pawn
			playerTour = true
		}

		// Check for a win or draw
		if checkWin(matrix, firstPlayer.Pawn) {
			fmt.Println("Le joueur", firstPlayer.Pseudo, "a gagné !")
			break
		} else if checkWin(matrix, secondPlayer.Pawn) {
			fmt.Println("Le joueur", secondPlayer.Pseudo, "a gagné !")
			break
		} else if isDraw(matrix) {
			fmt.Println("Match nul !")
			break
		} else {
			fmt.Println("Le jeu continue...")
		}
	}
}

// Function to check if a player has won
func checkWin(matrix [][]string, pawn string) bool {
	// Check rows and columns
	for i := 0; i < len(matrix); i++ {
		rowOccupied, colOccupied := true, true
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != pawn {
				rowOccupied = false
			}
			if matrix[j][i] != pawn {
				colOccupied = false
			}
		}
		if rowOccupied || colOccupied {
			return true
		}
	}

	// Check diagonals
	diagonalOccupied1, diagonalOccupied2 := true, true
	for i := 0; i < len(matrix); i++ {
		if matrix[i][i] != pawn {
			diagonalOccupied1 = false
		}
		if matrix[i][len(matrix)-1-i] != pawn {
			diagonalOccupied2 = false
		}
	}
	if diagonalOccupied1 || diagonalOccupied2 {
		return true
	}

	return false
}

// Function to check if the game is a draw
func isDraw(matrix [][]string) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "◻️ " {
				return false
			}
		}
	}
	return true
}
