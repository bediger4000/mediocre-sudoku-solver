package board

import "fmt"

var buddyThing []ThingType = []ThingType{RowThing, ColumnThing, BlockThing}

func (bd *Board) BlockOnly(announceSolution bool) int {
	found := 0
	for neighborhoodNo, neighborhood := range Neighborhoods {
		for buddiesNo, buddies := range neighborhood {
			digitCount := bd.CountPossibleDigits(buddyThing[neighborhoodNo], buddiesNo)
			for _, cell := range buddies {
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
	}
	return found
}
