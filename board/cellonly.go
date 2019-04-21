package board

func (bd *Board) BlockOnly() int {
	found := 0
	for blockno, block := range Blocks {
		digitCount := bd.CountPossibleDigits(BlockThing, blockno)

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
