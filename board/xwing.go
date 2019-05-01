package board

import "fmt"

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
								fmt.Printf("%s %d, <%d,%d>/<%d,%d>: common digit %d, non-common digits %d & %d\n",
									things[neighborhoodNo].nName, buddyNumber, row, col, row2, col2, commonDigit, diffDigit1, diffDigit2)
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
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col, bd[r][col].Possible[0], bd[r][col].Possible[2])
											}
										}
										if len(bd[r][col2].Possible) == 2 {
											if bd[r][col2].Possible[0] == diffDigit1 && bd[r][col2].Possible[1] == diffDigit2 {
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col2, bd[r][col2].Possible[0], bd[r][col2].Possible[2])
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
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col, bd[r][col].Possible[0], bd[r][col].Possible[1])
											}
										}
										if len(bd[r][col2].Possible) == 2 {
											if bd[r][col2].Possible[0] == diffDigit1 && bd[r][col2].Possible[1] == diffDigit2 {
												fmt.Printf("  <%d,%d> contains the non-common digits %d and %d\n", r, col2, bd[r][col2].Possible[0], bd[r][col2].Possible[1])
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
