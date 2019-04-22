package main

import (
	"fmt"
	"log"
	"os"
	"sudoku/board"
)

func main() {
	f1, e := os.Open(os.Args[1])
	if e != nil {
		log.Fatal(e)
	}
	fmt.Fprintf(os.Stderr, "Opened %q\n", os.Args[1])
	f2, e := os.Open(os.Args[2])
	if e != nil {
		log.Fatal(e)
	}
	fmt.Fprintf(os.Stderr, "Opened %q\n", os.Args[2])

	bd1 := board.ReadBoard(f1)
	fmt.Fprintf(os.Stderr, "Read in board 1\n")
	bd2 := board.ReadBoard(f2)
	fmt.Fprintf(os.Stderr, "Read in board 2\n")
	f1.Close()
	f2.Close()

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if bd1[row][col].Solved && bd2[row][col].Solved {
				if bd1[row][col].Value != bd2[row][col].Value {
					fmt.Printf("<%d,%d>, board 1 %d != board 2 %d\n", row, col, bd1[row][col].Value, bd2[row][col].Value)
				}
			}
		}
	}

}
