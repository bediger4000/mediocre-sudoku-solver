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
	var doNakedSubset bool
	var doHiddenPair bool
	var doXYWing bool
	var validateOnly bool
	flag.BoolVar(&printPossible, "c", false, "on incomplete solution, print digit possibilities")
	flag.BoolVar(&printPossibleExit, "C", false, "read input board, print digit possibilities, exit")
	flag.BoolVar(&announceSolution, "a", false, "announce solutions of cells")
	flag.BoolVar(&testingOutput, "f", false, "final solution output only")
	flag.BoolVar(&doNakedSubset, "N", false, "perform naked subset solving")
	flag.BoolVar(&doHiddenPair, "H", false, "perform hidden pair elimination")
	flag.BoolVar(&doXYWing, "X", false, "perform XY-wing elimination")
	flag.BoolVar(&validateOnly, "v", false, "validate the input board, then exit")
	psOutputNamePtr := flag.String("p", "", "PostScript output file name")
	flag.Parse()

	if testingOutput {
		announceSolution = false
	}

	fin := os.Stdin
	if flag.NArg() > 0 {
		var err error
		fin, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	bd := board.ReadBoard(fin)
	board.ValidateCells()

	if validateOnly {
		var sentence string
		if bd.IncompleteSolution() {
			phrase := "Incomplete Solution, "
			suffix := "invalid\n"
			if bd.CheckIntermediateValidity() {
				suffix = "valid\n"
			}
			sentence = phrase + suffix
		} else {
			phrase := "Complete Solution, "
			suffix := "invalid\n"
			if bd.CheckValidity() {
				suffix = "valid\n"
			}
			sentence = phrase + suffix
		}
		fmt.Printf(sentence)
		return
	}

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

		if doHiddenPair {
			totalFilled += bd.HiddenPair(announceSolution)
		}

		if doNakedSubset {
			totalFilled += bd.NakedSubset(announceSolution)
		}

		if doXYWing {
			totalFilled += bd.XYwing(announceSolution)
		}

		if !testingOutput {
			fmt.Printf("Filled or elimated %d cells\n", totalFilled)
			fmt.Printf(">>intermediate==\n")
			bd.PrintBoard(os.Stdout)
		}
	}
	if !testingOutput {
		fmt.Printf("==final==\n")
	}
	bd.PrintBoard(os.Stdout)
	if !testingOutput {
		if bd.IncompleteSolution() {
			fmt.Printf("===Incomplete Solution===\n")
			if !bd.CheckIntermediateValidity() {
				fmt.Printf("!!! Invalid !!!\n")
			}
		} else {
			fmt.Printf("===Complete Solution===\n")
			if !bd.CheckValidity() {
				fmt.Printf("!!! Invalid !!!\n")
			}
		}
	}
}
