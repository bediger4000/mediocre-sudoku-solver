package board

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (bd *Board) PrintBoard(out io.Writer) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if bd[i][j].Solved {
				fmt.Fprintf(out, "%1d ", bd[i][j].Value)
			} else {
				fmt.Fprintf(out, "_ ")
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func ReadBoard(in io.Reader) *Board {
	bd := NewBoard()
	r := bufio.NewReader(os.Stdin)
	for row := 0; row < 9; {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if buf[0] == '#' {
			continue
		}
		buf = bytes.Trim(buf, " \t\n")

		col := 0
		for _, c := range buf {
			if c == ',' || c == ' ' {
				continue
			}
			n := int(c - '0')
			if c == '_' {
				n = 0
			}
			if n < 0 || n > 10 {
				log.Fatalf("Numbers must be less than 10, greater than zero: %d (%c)\n", n, c)
			}
			if n != 0 {
				bd[row][col].Value = n
				bd[row][col].Solved = true
				bd[row][col].Possible = []int{n}
			}
			col++
		}
		row++
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				bd.EliminatePossibilities(row, col, bd[row][col].BlockNo, bd[row][col].Value)
			}
		}
	}
	return bd
}

func (bd *Board) PrintPossibilities(out io.Writer) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("<%d,%d> ", row, col)
			for _, digit := range bd[row][col].Possible {
				if digit != 0 {
					fmt.Printf("%d ", digit)
				}
			}
			fmt.Printf("\n")
		}
	}
}
