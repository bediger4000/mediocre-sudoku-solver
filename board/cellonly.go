package board

import "fmt"

func (bd *Board) BlockOnly(announceSolution bool) int {
	found := 0
	for blockno, block := range Blocks {
		digitCount := bd.CountPossibleDigits(BlockThing, blockno)

		for digit, count := range digitCount {
			if count == 1 {
				for _, cell := range block {
					for _, possibleDigit := range bd[cell.X][cell.Y].Possible {
						if possibleDigit == digit {
							if announceSolution {
								fmt.Printf("Mark <%d,%d> solved with %d, only possible digit for block\n", cell.X, cell.Y, digit)
							}
							bd.MarkSolved(cell.X, cell.Y, digit)
							found++
							break
						}
					}
				}
			}
		}
	}
	return found
}
