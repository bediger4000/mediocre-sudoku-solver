package board

import "fmt"

func (bd *Board) HiddenSubset(announceSolution bool) int {
	//fmt.Printf("Enter HiddenSubset\n")
	//defer fmt.Printf("Exit HiddenSubset\n")
	solved := 0
	for col := 0; col < 9; col++ {
		if u := bd.CountUnsolvedRows(col); u < 3 {
			//fmt.Printf("Col %d only has %d unsolved rows\n", col, u)
			continue
		} //else {
		//	fmt.Printf("Considering col %d for hidden subset, %d unsolved rows\n", col, u)
		//}

		triples := [][]int{}
		triplesRows := [][2]int{}
		for i := 0; i < 8; i++ {
			if bd[i][col].Solved {
				// fmt.Printf("->Col %d row %d solved with %d\n", col, i, bd[i][col].Value)
				continue
			}
			// fmt.Printf("- Comparing Col %d, row %d against subsequent rows for common digits\n", col, i)
			for j := i + 1; j < 9; j++ {
				if bd[j][col].Solved {
					// fmt.Printf("-->Col %d row %d solved with %d\n", col, j, bd[j][col].Value)
					continue
				}
				common := commonDigits(bd[i][col].Possible, bd[j][col].Possible)
				// fmt.Printf("Col %d, rows %d & %d common digits: %v\n", col, i, j, common)
				if len(common) == 3 {
					triples = append(triples, common)
					triplesRows = append(triplesRows, [2]int{i, j})
				}
			}
		}

		if len(triples) != 1 {
			// fmt.Printf("Column %d only has %d triples\n", col, len(triples))
			continue
		}

		for idx, triplet := range triples {
			r1 := triplesRows[idx][0]
			r2 := triplesRows[idx][1]
			// fmt.Printf("Col %d, rows %d,%d have %v in common\n", col, r1, r2, triplet)
			for _, digit := range triplet {
				singleDigits := []int{}
				singleDigitRows := []int{}
				for row := 0; row < 9; row++ {
					if row != r1 && row != r2 && !bd[row][col].Solved {
						for _, single := range bd[row][col].Possible {
							if single == 0 {
								continue
							}
							if single == digit {
								singleDigits = append(singleDigits, digit)
								singleDigitRows = append(singleDigitRows, row)

							}
						}
					}
				}
				if len(singleDigits) == 1 {
					//fmt.Printf("Column %d Row %d has a single %d in it, hidden subset with rows %d & %d\n",
					//	col, singleDigitRows[0], singleDigits[0], r1, r2)
					// Column col, row singleDigitRows[0] has solved value singleDigits[0].
					if announceSolution {
						fmt.Printf("Mark <%d,%d> solved with %d, hidden subset\n", singleDigitRows[0], col, singleDigits[0])
					}
					bd.MarkSolved(singleDigitRows[0], col, singleDigits[0])
					solved++
				}
			}
		}
	}
	return solved
}

func (bd *Board) CountUnsolvedRows(col int) int {
	unsolvedRows := 0
	for row := 0; row < 9; row++ {
		if !bd[row][col].Solved {
			unsolvedRows++
		}
	}
	return unsolvedRows
}

func commonDigits(poss1 []int, poss2 []int) []int {

	var common []int
	for _, digit := range poss1 {
		for _, otherdigit := range poss2 {
			if digit == otherdigit {
				common = append(common, digit)
			}
		}
	}

	return common
}
