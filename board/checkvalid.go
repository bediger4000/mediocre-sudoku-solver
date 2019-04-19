package board

import "fmt"

func (bd *Board) CheckIntermediateValidity() bool {
	for row := 0; row < 9; row++ {
		incomplete := false
		sum := 0
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for col := 0; col < 9; col++ {
			if !bd[row][col].Solved {
				incomplete = true
			}
			digitCounts[bd[row][col].Value]++
			sum += bd[row][col].Value
		}
		if !incomplete && sum != 45 {
			return false
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Row %d has %d %d digits\n", row, count, digit)
				return false
			}
		}
	}
	for col := 0; col < 9; col++ {
		incomplete := false
		sum := 0
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for row := 0; row < 9; row++ {
			if !bd[row][col].Solved {
				incomplete = true
			}
			digitCounts[bd[row][col].Value]++
			sum += bd[row][col].Value
		}
		if !incomplete && sum != 45 {
			return false
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Row %d has %d %d digits\n", col, count, digit)
				return false
			}
		}
	}
	for blockNo, block := range Blocks {
		incomplete := false
		sum := 0
		digitCounts := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, cell := range block {
			if !bd[cell.X][cell.Y].Solved {
				incomplete = true
			}
			sum += bd[cell.X][cell.Y].Value
			digitCounts[bd[cell.X][cell.Y].Value]++
		}
		if !incomplete && sum != 45 {
			return false
		}
		for digit, count := range digitCounts {
			if digit != 0 && count > 1 {
				fmt.Printf("Block %d has %d %d digits\n", blockNo, count, digit)
				return false
			}
		}
	}
	return true
}

func (bd *Board) CheckValidity() bool {
	valid := true
	for row := 0; row < 9; row++ {
		sum := 0
		for col := 0; col < 9; col++ {
			sum += bd[row][col].Value
		}
		if sum != 45 {
			valid = false
		}
	}
	for col := 0; col < 9; col++ {
		sum := 0
		for row := 0; row < 9; row++ {
			sum += bd[row][col].Value
		}
		if sum != 45 {
			valid = false
		}
	}

	for _, block := range Blocks {
		sum := 0
		for _, cell := range block {
			sum += bd[cell.X][cell.Y].Value
		}
		if sum != 45 {
			valid = false
		}
	}
	return valid
}
