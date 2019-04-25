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
	var announceSolution bool
	var testingOutput bool
	flag.BoolVar(&printPossible, "c", false, "on incomplete solution, print digit possibilities")
	flag.BoolVar(&printPossibleExit, "C", false, "read input board, print digit possibilities, exit")
	flag.BoolVar(&announceSolution, "a", false, "announce solutions of cells")
	flag.BoolVar(&testingOutput, "f", false, "final solution output only")
	psOutputNamePtr := flag.String("p", "", "PostScript output file name")
	flag.Parse()

	if testingOutput {
		announceSolution = false
	}

	bd := board.ReadBoard(os.Stdin)
	board.ValidateCells()

	if *psOutputNamePtr != "" {
		fd, err := os.Create(*psOutputNamePtr)
		if err != nil {
			log.Fatalf("opening %q: %v", *psOutputNamePtr, err)
		}
		bd.EmitPostScript(fd, printPossibleExit)
		fd.Close()
		return
	}

	if printPossibleExit {
		bd.PrintPossibilities(os.Stdout)
		return
	}

	totalFilled := 1
	for totalFilled > 0 {
		totalFilled = 0
		n := 1
		for n > 0 {
			n = bd.OnlyPossibility(announceSolution)
			totalFilled += n
		}

		n = 1
		for n > 0 {
			n = bd.BlockOnly(announceSolution)
			totalFilled += n
		}

		n = 1
		for n > 0 {
			n = bd.HiddenSubset(announceSolution)
			totalFilled += n
		}

		if !testingOutput {
			fmt.Printf("Filled in %d cells\n", totalFilled)
			fmt.Printf(">>intermediate==\n")
			bd.PrintBoard(os.Stdout)
		}

		if totalFilled == 0 {
			bd.NakedSubset()
		}
	}
	if !testingOutput {
		fmt.Printf("==final==\n")
	}
	bd.PrintBoard(os.Stdout)
	if !testingOutput {
		if bd.IncompleteSolution() {
			fmt.Printf("===Incomplete Solution===\n")
			bd.CheckIntermediateValidity()
		} else {
			fmt.Printf("===Complete Solution===\n")
			bd.CheckValidity()
		}
	}
}
