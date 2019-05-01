package board

import "fmt"

func (bd *Board) OnlyPossibility(announceSolution bool) int {
	found := 0
	for _, row := range Rows {
		for _, cell := range row {
			row, col := cell.X, cell.Y
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
