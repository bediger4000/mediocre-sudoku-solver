package main

import (
	"fmt"
	"log"
	"os"
	"sudoku/board"
)

func main() {

	bd := board.ReadBoard(os.Stdin)
	bd.XYwing(true)
	// fmt.Printf("Eliminated %d possible digits\n", eliminated)
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
