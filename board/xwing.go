package board

import "fmt"

// XYwing disgusts me. Really deeply nested.
func (bd *Board) XYwing(announce bool) int {
	eliminated := 0
	// Loop over Rows, Columns, Blocks
	for neighborhoodNo, neighborhood := range Neighborhoods {
		for buddyNumber, buddies := range neighborhood {
			for cellIdx, cell := range buddies {
				row, col := cell.X, cell.Y
				if bd[row][col].Solved {
					continue
				}
				if len(bd[row][col].Possible) == 2 {
					poss1 := bd[row][col].Possible
					for cellIdx2, cell2 := range buddies {
						if cellIdx > cellIdx2 {
							continue
						}
						row2, col2 := cell2.X, cell2.Y
						if row == row2 && col == col2 {
							continue
						}
						if bd[row2][col2].Solved {
							continue
						}
						if len(bd[row2][col2].Possible) == 2 {
							poss2 := bd[row2][col2].Possible
							commonDigit, diffDigit1, diffDigit2, found := checkCommon(poss1, poss2)
							if found {
								// a common digit, and two non-common digits
								// Find a non-common digits in a column, if neighborhoodNo == 0 (Rows),
								// find a non-common digits in a row, if neighborhoodNo == 1 (Columns),
								// find a non-common digits in a row or a colum, if neighborhoodNo == 2 (Blocks)
								if things[neighborhoodNo].nType == RowThing {
									// row == row2, so look in col and col2
									for r := 0; r < 9; r++ {
										if r == row {
											continue
										}

										if len(bd[r][col].Possible) == 2 {
											if bd[r][col].Possible[0] == diffDigit1 && bd[r][col].Possible[1] == diffDigit2 {
												fmt.Printf("%s %d, <%d,%d>/<%d,%d>: common digit %d, non-common digits %d & %d\n",
													things[neighborhoodNo].nName, buddyNumber, row, col, row2, col2, commonDigit, diffDigit1, diffDigit2)
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col, bd[r][col].Possible[0], bd[r][col].Possible[1])
												// bd[row][col]: wings bd[row2][col2], bd[r][col]
											}
										}

										if len(bd[r][col2].Possible) == 2 {
											if bd[r][col2].Possible[0] == diffDigit1 && bd[r][col2].Possible[1] == diffDigit2 {
												fmt.Printf("%s %d, <%d,%d>/<%d,%d>: common digit %d, non-common digits %d & %d\n",
													things[neighborhoodNo].nName, buddyNumber, row, col, row2, col2, commonDigit, diffDigit1, diffDigit2)
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col2, bd[r][col2].Possible[0], bd[r][col2].Possible[1])
												// bd[row][col]: wings bd[row2][col2], bd[r][col2]
											}
										}
									}
								}
								if things[neighborhoodNo].nType == ColumnThing {
									// col == col2, look in row and row2
								}
								if things[neighborhoodNo].nType == BlockThing {
									// row not necessarily == row2, col not necessarily == col2
									// <row,col> and <row2,col2> have commonDigit in common, diffDigit1, diffDigit2 non-common
									// Look up and down col and col2 for cell with only diffDigit1, diffDigit2 as possible
									for r := 0; r < 9; r++ {
										if r == row {
											continue
										}
										if len(bd[r][col].Possible) == 2 {
											if bd[r][col].Possible[0] == diffDigit1 && bd[r][col].Possible[1] == diffDigit2 {
												if announce {
													fmt.Printf("%s %d, <%d,%d>/<%d,%d>: common digit %d, non-common digits %d & %d\n",
														things[neighborhoodNo].nName, buddyNumber, row, col, row2, col2, commonDigit, diffDigit1, diffDigit2)
													fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col, bd[r][col].Possible[0], bd[r][col].Possible[1])
												}
												eliminateDigit, _, _, foundEliminate := checkCommon(bd[row2][col2].Possible, bd[r][col].Possible)
												if foundEliminate {
													if announce {
														fmt.Printf("  Could eliminate %d from any cells visible to both <%d,%d> & <%d,%d>\n",
															eliminateDigit, row2, col2, r, col)
													}
													eliminated += bd.eliminateMutuallyVisible(eliminateDigit, row2, col2, r, col, announce)
												}
											}
										}
										if len(bd[r][col2].Possible) == 2 {
											if bd[r][col2].Possible[0] == diffDigit1 && bd[r][col2].Possible[1] == diffDigit2 {
												if announce {
													fmt.Printf("%s %d, <%d,%d>/<%d,%d>: common digit %d, non-common digits %d & %d\n",
														things[neighborhoodNo].nName, buddyNumber, row, col, row2, col2, commonDigit, diffDigit1, diffDigit2)
													fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col2, bd[r][col2].Possible[0], bd[r][col2].Possible[1])
												}
												eliminateDigit, _, _, foundEliminate := checkCommon(bd[row2][col2].Possible, bd[r][col2].Possible)
												if foundEliminate {
													if announce {
														fmt.Printf("  Could eliminate %d from any cells visible to both <%d,%d> & <%d,%d>\n",
															eliminateDigit, row2, col2, r, col2)
													}
													eliminated += bd.eliminateMutuallyVisible(eliminateDigit, row2, col2, r, col2, announce)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return eliminated
}

// noncommon1 < noncommon2
func checkCommon(poss1 []int, poss2 []int) (commonDigit int, noncommon1 int, noncommon2 int, foundThem bool) {
	if poss1[0] == poss2[0] {
		commonDigit, noncommon1, noncommon2, foundThem = poss1[0], poss1[1], poss2[1], true
	}
	if poss1[0] == poss2[1] {
		commonDigit, noncommon1, noncommon2, foundThem = poss1[0], poss1[1], poss2[0], true
	}
	if poss1[1] == poss2[0] {
		commonDigit, noncommon1, noncommon2, foundThem = poss1[1], poss1[0], poss2[1], true
	}
	if poss1[1] == poss2[1] {
		commonDigit, noncommon1, noncommon2, foundThem = poss1[1], poss1[0], poss2[0], true
	}
	if noncommon1 > noncommon2 {
		noncommon1, noncommon2 = noncommon2, noncommon1
	}
	if noncommon1 == noncommon2 {
		foundThem = false
	}
	return commonDigit, noncommon1, noncommon2, foundThem
}

func (bd *Board) eliminateMutuallyVisible(digit, x0, y0, x1, y1 int, announce bool) int {
	eliminated := 0
	// if bd[x0][y0].BlockNo == bd[x1][y1].BlockNo {
	if announce {
		fmt.Printf("Elminate %d from <%d,%d> & <%d,%d> mutuals\n", digit, x0, y0, x1, y1)
	}
	if !bd[x0][y1].Solved {
		if bd.SpliceOut(x0, y1, digit) > 0 {
			eliminated++
			if announce {
				fmt.Printf("Eliminated %d from <%d,%d> possibles\n", digit, x0, y1)
			}
		}
	}
	if !bd[x1][y0].Solved {
		if bd.SpliceOut(x1, y0, digit) > 0 {
			eliminated++
			if announce {
				fmt.Printf("Eliminated %d from <%d,%d> possibles\n", digit, x1, y0)
			}
		}
	}
	if bd[x0][y0].BlockNo == bd[x1][y1].BlockNo {
		for _, cell := range Blocks[bd[x0][y0].BlockNo] {
			if cell.X == x0 && cell.Y == y0 || cell.X == x1 && cell.Y == y1 {
				continue
			}
			if !bd[cell.X][cell.Y].Solved {
				if bd.SpliceOut(cell.X, cell.Y, digit) > 0 {
					eliminated++
					if announce {
						fmt.Printf("Eliminated %d from <%d,%d> possibles\n", digit, cell.X, cell.Y)
					}
				}
			}
		}
	}
	return eliminated
}
