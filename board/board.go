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

func (bd *Board) SpliceOut(row, col, digitEliminate int) {
	for idx, digit := range bd[row][col].Possible {
		if digit == digitEliminate {
			bd[row][col].Possible = append(bd[row][col].Possible[:idx], bd[row][col].Possible[idx+1:]...)
			break
		}
	}
}

func (bd *Board) EliminatePossibilities(rowEliminate, colEliminate, blockEliminate, digitEliminate int) {
	for col := 0; col < 9; col++ {
		if bd[rowEliminate][col].Solved {
			continue
		}
		if col == colEliminate {
			continue
		}
		bd.SpliceOut(rowEliminate, col, digitEliminate)
	}
	for row := 0; row < 9; row++ {
		if bd[row][colEliminate].Solved {
			continue
		}
		if row == rowEliminate {
			continue
		}
		bd.SpliceOut(row, colEliminate, digitEliminate)
	}

	for _, coord := range Blocks[blockEliminate] {
		if coord.X == rowEliminate && coord.Y == colEliminate {
			continue
		}
		if bd[coord.X][coord.Y].Solved {
			continue
		}
		bd.SpliceOut(coord.X, coord.Y, digitEliminate)
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
