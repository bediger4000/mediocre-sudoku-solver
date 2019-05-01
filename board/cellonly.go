package board

import "fmt"

func (bd *Board) BlockOnly(announceSolution bool) int {
	found := 0
	for blockno, block := range Blocks {
		digitCount := bd.CountPossibleDigits(BlockThing, blockno)

		for _, cell := range block {
			row, col := cell.X, cell.Y
			if bd[row][col].Solved {
				continue
			}
			for _, possibleDigit := range bd[row][col].Possible {
				if digitCount[possibleDigit] == 1 {
					if announceSolution {
						fmt.Printf("Mark <%d,%d> solved with %d, only possible digit for block\n", row, col, possibleDigit)
					}
					bd.MarkSolved(row, col, possibleDigit)
					found++
					break
				}
			}
		}
	}
	return found
}
