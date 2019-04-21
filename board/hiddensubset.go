package board

import "fmt"

func (bd *Board) HiddenSubset() int {
	fmt.Printf("Enter HiddenSubset\n")
	defer fmt.Printf("Exit HiddenSubset\n")
	for col := 0; col < 9; col++ {
		if bd.CountUnsolvedRows(col) < 3 {
			continue
		}
		digitCounts := bd.CountPossibleDigits(ColumnThing, col)
		l, m, n, doesNot := mightHaveHiddenSubset(digitCounts)
		if doesNot {
			fmt.Printf("Col %d does not\n", col)
			continue
		}
		// Find only 2 cells with l,m,n as possible, just one other cell with n as possible
		all3Digits := []int{} // row numbers
		just1Digit := 0       // row number
		for row := 0; row < 9; row++ {
			count := countDigits(l, m, n, bd[row][col].Possible)
			switch count {
			case 3:
				all3Digits = append(all3Digits, row)
			case 1:
				just1Digit = row
			}
		}
		if len(all3Digits) != 2 {
			continue
		}
		// col,{all3Digits...} can only have l,m in them
		// col,just1Digit has n in it
		fmt.Printf("Col %d has hidden subset [%d,%d,%d] at rows %v, %d out\n",
			col, l, m, n, all3Digits, just1Digit,
		)
	}
	return 0
}

// 2 cells with the same triplet, and only one cell that has
// one digit of the triple, the other 2 digits only in the 2 cells
// 2 digits with a 2-count, one digit with a 3-count
func mightHaveHiddenSubset(digitCounts [10]int) (int, int, int, bool) {
	fmt.Printf("Enter mightHaveHiddenSubsets %v\n", digitCounts)
	defer fmt.Printf("Exit mightHaveHiddenSubsets\n")
	twoCounts := 0
	threeCounts := 0
	twoCountDigits := make([]int, 0)
	threeCountDigit := 0
	for digit, count := range digitCounts {
		switch count {
		case 2:
			twoCountDigits = append(twoCountDigits, digit)
			twoCounts++
		case 3:
			threeCountDigit = digit
			threeCounts++
		}
	}
	if twoCounts != 2 || threeCounts != 1 {
		fmt.Printf("Not a HiddenSubset 2-counts: %d, 3-counts: %d\n", twoCounts, threeCounts)
		return 0, 0, 0, false
	}

	fmt.Printf("maybe a HiddenSubset [%d,%d,%d]\n", twoCountDigits[0], twoCountDigits[1], threeCountDigit)
	return twoCountDigits[0], twoCountDigits[1], threeCountDigit, true
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

func countDigits(l, m, n int, ary []int) int {
	count := 0
	for _, x := range ary {
		switch {
		case l == x:
			count++
		case m == x:
			count++
		case n == x:
			count++
		}
	}
	return count
}
