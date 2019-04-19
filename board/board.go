package board

import "fmt"

func ValidateCells() {
	for cellno, cell := range Cells {
		cellIndex := cell.X*9 + cell.Y
		if cellno != cellIndex {
			fmt.Printf("Cell number %d at <%d,%d> has cell index %d\n", cellno, cell.X, cell.Y, cellIndex)
		}
		inTheBlock := false
		block := Blocks[cell.BlockNo]
		for _, aCell := range block {
			if cell.X == aCell.X && cell.Y == aCell.Y {
				inTheBlock = true
				break
			}
		}
		if !inTheBlock {
			fmt.Printf("Could not find cell %d in block %d\n", cellno, cell.BlockNo)
		}
	}
}

func NewBoard() *Board {
	var bd Board

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			bd[row][col].X = row
			bd[row][col].Y = col
			bd[row][col].Solved = false
			bd[row][col].Value = 0
			Cells = append(Cells, &bd[row][col])
			bd[row][col].Possible = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}

	for blockno, block := range Blocks {
		for _, cell := range block {
			bd[cell.X][cell.Y].BlockNo = blockno
		}
	}

	return &bd
}

func (bd *Board) MarkSolved(row, col int, digit int) {
	bd[row][col].Value = digit
	bd[row][col].Possible = []int{}
	bd[row][col].Solved = true
	bd.EliminatePossibilities(row, col, bd[row][col].BlockNo, bd[row][col].Value)
}

func (bd *Board) EliminatePossibilities(rowEliminate, colEliminate, blockEliminate, digitEliminate int) {
	for col := 0; col < 9; col++ {
		if bd[rowEliminate][col].Solved {
			continue
		}
		if col == colEliminate {
			continue
		}
		for idx, digit := range bd[rowEliminate][col].Possible {
			if digit == digitEliminate {
				bd[rowEliminate][col].Possible = append(bd[rowEliminate][col].Possible[:idx], bd[rowEliminate][col].Possible[idx+1:]...)
				break
			}
		}
	}
	for row := 0; row < 9; row++ {
		if bd[row][colEliminate].Solved {
			continue
		}
		if row == rowEliminate {
			continue
		}
		for idx, digit := range bd[row][colEliminate].Possible {
			if digit == digitEliminate {
				bd[row][colEliminate].Possible = append(bd[row][colEliminate].Possible[:idx], bd[row][colEliminate].Possible[idx+1:]...)
				break
			}
		}
	}

	for _, coord := range Blocks[blockEliminate] {
		if coord.X == rowEliminate && coord.Y == colEliminate {
			continue
		}
		if bd[coord.X][coord.Y].Solved {
			continue
		}
		for idx, digit := range bd[coord.X][coord.Y].Possible {
			if digit == digitEliminate {
				bd[coord.X][coord.Y].Possible = append(bd[coord.X][coord.Y].Possible[:idx], bd[coord.X][coord.Y].Possible[idx+1:]...)
				break
			}
		}
	}
}

func (bd *Board) IncompleteSolution() bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if !bd[row][col].Solved {
				return true
			}
		}
	}
	return false
}
