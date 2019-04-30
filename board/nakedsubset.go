package board

import "fmt"

func (bd *Board) NakedSubset(announce bool) int {

	eliminated := 0

	for row := 0; row < 9; row++ {
		digitCount := bd.CountPossibleDigits(RowThing, row)
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
				if digitCount[bd[row][twoCols[0]].Possible[0]] > 2 &&
					digitCount[bd[row][twoCols[0]].Possible[1]] > 2 {
					continue
				}
				if announce {
					fmt.Printf("Row %d, naked subset of [%d %d] in <%d,%d> and <%d,%d>\n",
						row,
						bd[row][twoCols[0]].Possible[0], bd[row][twoCols[0]].Possible[1],
						row, twoCols[0], row, twoCols[1],
					)
				}
				for c := 0; c < 9; c++ {
					if bd[row][c].Solved {
						continue
					}
					if c == twoCols[0] || c == twoCols[1] {
						continue
					}
					// Splice either of the two digits out of bd[row][c].Possible
					n := bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[0])
					if announce && n > 0 {
						fmt.Printf("Row %d, spliced %d from <%d,%d>\n", row, bd[row][twoCols[0]].Possible[0], row, c)
					}
					m := bd.SpliceOut(row, c, bd[row][twoCols[0]].Possible[1])
					if announce && m > 0 {
						fmt.Printf("Row %d, spliced %d from <%d,%d>\n", row, bd[row][twoCols[0]].Possible[1], row, c)
					}
					eliminated += m + n
				}
			}
		}
	}

	for col := 0; col < 9; col++ {
		twoCounts := 0
		digitCount := bd.CountPossibleDigits(ColumnThing, col)
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
				if digitCount[bd[twoRows[0]][col].Possible[0]] > 2 &&
					digitCount[bd[twoRows[0]][col].Possible[1]] > 2 {
					continue
				}
				if announce {
					fmt.Printf("Col %d, naked subset of [%d %d] in <%d,%d> and <%d,%d>\n",
						col,
						bd[twoRows[0]][col].Possible[0], bd[twoRows[0]][col].Possible[1],
						twoRows[0], col, twoRows[1], col,
					)
				}
				for r := 0; r < 9; r++ {
					if bd[r][col].Solved {
						continue
					}
					if r == twoRows[0] || r == twoRows[1] {
						continue
					}
					// Splice either of the two digits out of bd[row][c].Possible
					n := bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[0])
					if announce && n > 0 {
						fmt.Printf("Col %d, spliced %d from <%d,%d>\n", col, bd[twoRows[0]][col].Possible[0], r, col)
					}
					m := bd.SpliceOut(r, col, bd[twoRows[0]][col].Possible[1])
					if announce && m > 0 {
						fmt.Printf("Col %d, spliced %d from <%d,%d>\n", col, bd[twoRows[0]][col].Possible[1], r, col)
					}
					eliminated += n + m
				}
			}
		}
	}
	for blockNo, block := range Blocks {
		twoCounts := 0
		digitCount := bd.CountPossibleDigits(BlockThing, blockNo)
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
				if digitCount[bd[X0][Y0].Possible[0]] > 2 &&
					digitCount[bd[X0][Y0].Possible[1]] > 2 {
					continue
				}
				if announce {
					fmt.Printf("Block %d, naked subset of [%d %d] in <%d,%d> and <%d,%d>\n",
						blockNo,
						bd[X0][Y0].Possible[0], bd[X0][Y0].Possible[1],
						X0, Y0, X1, Y1,
					)
				}
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
					n := bd.SpliceOut(cell.X, cell.Y, bd[X0][Y0].Possible[0])
					if announce && n > 0 {
						fmt.Printf("Block %d, spliced %d from <%d,%d>\n", blockNo, bd[cell.X][cell.Y].Possible[0], cell.X, cell.Y)
					}
					m := bd.SpliceOut(cell.X, cell.Y, bd[X0][Y0].Possible[1])
					if announce && m > 0 {
						fmt.Printf("Block %d, spliced %d from <%d,%d>\n", blockNo, bd[cell.X][cell.Y].Possible[1], cell.X, cell.Y)
					}
					eliminated += n + m
				}
			}
		}
	}

	return eliminated
}
