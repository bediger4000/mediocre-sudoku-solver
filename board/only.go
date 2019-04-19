package board

func (bd *Board) OnlyPossibility() int {
	found := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if !bd[row][col].Solved {
				if len(bd[row][col].Possible) == 1 {
					bd.MarkSolved(row, col, bd[row][col].Possible[0])
					found++
				}
			}
		}
	}
	return found
}
