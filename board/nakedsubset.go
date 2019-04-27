package board

func (bd *Board) NakedSubset() bool {

	var foundSome bool

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
					foundSome = foundSome || bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[0])
					foundSome = foundSome || bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[1])
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
					foundSome = foundSome || bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[0])
					foundSome = foundSome || bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[1])
				}
			}
		}
	}
	for _, block := range Blocks {
		twoCounts := 0
		for _, cell := range block {
			if bd[cell.X][cell.Y].Solved {
				continue
			}
			if len(bd[cell.X][cell.Y].Possible) == 2 {
				twoCounts++
			}
		}
		if twoCounts != 2 {
			continue
		}
		twoCells := []CellCoord{}
		for _, cell := range block {
			if bd[cell.X][cell.Y].Solved {
				continue
			}
			if len(bd[cell.X][cell.Y].Possible) == 2 {
				twoCells = append(twoCells, cell)
			}
		}

		X0, Y0 := twoCells[0].X, twoCells[0].Y
		X1, Y1 := twoCells[1].X, twoCells[1].Y

		if bd[X0][Y0].Possible[0] == bd[X1][Y1].Possible[0] {
			if bd[X0][Y0].Possible[1] == bd[X1][Y1].Possible[1] {
				for _, cell := range block {
					if bd[cell.X][cell.Y].Solved {
						continue
					}
					if cell.X == X0 && cell.Y == Y0 {
						continue
					}
					if cell.X == X1 && cell.Y == Y1 {
						continue
					}
					// Splice either of the two digits out of bd[row][c].Possible
					foundSome = foundSome || bd.SpliceOut(cell.X, cell.Y, bd[X0][Y0].Possible[0])
					foundSome = foundSome || bd.SpliceOut(cell.X, cell.Y, bd[X0][Y0].Possible[1])
				}
			}
		}
	}

	return foundSome
}
