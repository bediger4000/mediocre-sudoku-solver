package main

import (
	"os"
	"sudoku/board"
)

func main() {

	bd := board.ReadBoard(os.Stdin)
	bd.XYwing(true)
	// fmt.Printf("Eliminated %d possible digits\n", eliminated)
}
