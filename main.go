package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sudoku/board"
)

func main() {
	var printPossible bool
	var printPossibleExit bool
	flag.BoolVar(&printPossible, "c", false, "on incomplete solution, print digit possibilities")
	flag.BoolVar(&printPossibleExit, "C", false, "read input board, print digit possibilities, exit")
	psOutputNamePtr := flag.String("p", "", "PostScript output file name")
	flag.Parse()

	bd := board.ReadBoard(os.Stdin)

	if *psOutputNamePtr != "" {
		fd, err := os.Create(*psOutputNamePtr)
		if err != nil {
			log.Fatalf("opening %q: %v", *psOutputNamePtr, err)
		}
		bd.EmitPostScript(fd)
		fd.Close()
		return
	}

	if printPossibleExit {
		bd.PrintPossibilities(os.Stdout)
		return
	}

	n := bd.OnlyPossibility()
	fmt.Printf("Filled in %d cells\n", n)
	board.ValidateCells()
}
