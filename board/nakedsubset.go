package board

func (bd *Board) NakedSubset() {

	for row := 0; row < 9; row++ {
		twoCounts := 0
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				continue
			}
			if len(bd[row][col].Possible) == 2 {
				twoCounts++
			}
		}
		if twoCounts != 2 {
			continue
		}
		twoCols := []int{}
		for col := 0; col < 9; col++ {
			if bd[row][col].Solved {
				continue
			}
			if len(bd[row][col].Possible) == 2 {
				twoCols = append(twoCols, col)
			}
		}
		if bd[row][twoCols[0]].Possible[0] == bd[row][twoCols[1]].Possible[0] {
			if bd[row][twoCols[0]].Possible[1] == bd[row][twoCols[1]].Possible[1] {
				for c := 0; c < 9; c++ {
					if bd[row][c].Solved {
						continue
					}
					if c == twoCols[0] || c == twoCols[1] {
						continue
					}
					// Splice either of the two digits out of bd[row][c].Possible
					bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[0])
					bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[1])
				}
			}
		}
	}

	for col := 0; col < 9; col++ {
		twoCounts := 0
		for row := 0; row < 9; row++ {
			if bd[row][col].Solved {
				continue
			}
			if len(bd[row][col].Possible) == 2 {
				twoCounts++
			}
		}
		if twoCounts != 2 {
			continue
		}
		twoRows := []int{}
		for row := 0; row < 9; row++ {
			if bd[row][col].Solved {
				continue
			}
			if len(bd[row][col].Possible) == 2 {
				twoRows = append(twoRows, row)
			}
		}
		if bd[twoRows[0]][col].Possible[0] == bd[twoRows[1]][col].Possible[0] {
			if bd[twoRows[0]][col].Possible[1] == bd[twoRows[1]][col].Possible[1] {
				for r := 0; r < 9; r++ {
					if bd[r][col].Solved {
						continue
					}
					if r == twoRows[0] || r == twoRows[1] {
						continue
					}
					// Splice either of the two digits out of bd[row][c].Possible
					bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[0])
					bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[1])
				}
			}
		}
	}
}
