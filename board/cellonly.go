package board

func (bd *Board) BlockOnly() int {
	found := 0
	for _, block := range Blocks {
		digitCount := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for _, cell := range block {
			for _, digit := range bd[cell.X][cell.Y].Possible {
				digitCount[digit]++
			}
		}
		for digit, count := range digitCount {
			if count == 1 {
				for _, cell := range block {
					for _, possibleDigit := range bd[cell.X][cell.Y].Possible {
						if possibleDigit == digit {
							bd.MarkSolved(cell.X, cell.Y, digit)
							found++
							break
						}
					}
				}
			}
		}
	}
	return found
}
