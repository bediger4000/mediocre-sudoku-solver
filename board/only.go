package board

import "fmt"

func (bd *Board) OnlyPossibility(announceSolution bool) int {
	found := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if !bd[row][col].Solved {
				if len(bd[row][col].Possible) == 1 {
					if announceSolution {
						fmt.Printf("Mark <%d,%d> solved with %d, sole possibility\n", row, col, bd[row][col].Possible[0])
					}
					bd.MarkSolved(row, col, bd[row][col].Possible[0])
					found++
				}
			}
		}
	}
	return found
}
