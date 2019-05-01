package main

import (
	"fmt"
	"log"
	"os"
	"sudoku/board"
)

func main() {

	bd := board.ReadBoard(os.Stdin)
	board.ValidateCells()

	n := bd.HiddenPair(true)
	fmt.Printf("Found %d hidden subsets\n", n)

	if len(os.Args) > 1 {
		psFileName := os.Args[1]
		fmt.Printf("Putting PS with possibilities in %q\n", psFileName)
		fd, err := os.Create(psFileName)
		if err != nil {
			log.Fatalf("opening %q: %v", psFileName, err)
		}
		bd.EmitPostScript(fd, true)
		fd.Close()
	}
}
