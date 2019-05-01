package main

import (
	"fmt"
	"os"
	"sudoku/board"
)

func main() {

	bd := board.ReadBoard(os.Stdin)

	eliminated := bd.NakedSubset(true)

	fmt.Printf("Eliminated %d possible digits\n", eliminated)

	if len(os.Args) > 1 {
		psFileName := os.Args[1]
		fd, e := os.Create(psFileName)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Not creating PS file %q: %v\n", psFileName, e)
		}
		bd.EmitPostScript(fd, true)
		fd.Close()
	}
}
