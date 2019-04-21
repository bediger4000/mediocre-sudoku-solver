package board

type Possible []int

type Cell struct {
	X        int
	Y        int
	BlockNo  int
	Solved   bool
	Value    int
	Possible []int
}

type Board [9][9]Cell

type CellCoord struct {
	X int
	Y int
}

type Block [9]CellCoord

var Blocks [9]Block = [9]Block{
	Block{CellCoord{X: 0, Y: 0}, CellCoord{X: 0, Y: 1}, CellCoord{X: 0, Y: 2}, CellCoord{X: 1, Y: 0}, CellCoord{X: 1, Y: 1}, CellCoord{X: 1, Y: 2}, CellCoord{X: 2, Y: 0}, CellCoord{X: 2, Y: 1}, CellCoord{X: 2, Y: 2}},
	Block{CellCoord{X: 0, Y: 3}, CellCoord{X: 0, Y: 4}, CellCoord{X: 0, Y: 5}, CellCoord{X: 1, Y: 3}, CellCoord{X: 1, Y: 4}, CellCoord{X: 1, Y: 5}, CellCoord{X: 2, Y: 3}, CellCoord{X: 2, Y: 4}, CellCoord{X: 2, Y: 5}},
	Block{CellCoord{X: 0, Y: 6}, CellCoord{X: 0, Y: 7}, CellCoord{X: 0, Y: 8}, CellCoord{X: 1, Y: 6}, CellCoord{X: 1, Y: 7}, CellCoord{X: 1, Y: 8}, CellCoord{X: 2, Y: 6}, CellCoord{X: 2, Y: 7}, CellCoord{X: 2, Y: 8}},
	Block{CellCoord{X: 3, Y: 0}, CellCoord{X: 3, Y: 1}, CellCoord{X: 3, Y: 2}, CellCoord{X: 4, Y: 0}, CellCoord{X: 4, Y: 1}, CellCoord{X: 4, Y: 2}, CellCoord{X: 5, Y: 0}, CellCoord{X: 5, Y: 1}, CellCoord{X: 5, Y: 2}},
	Block{CellCoord{X: 3, Y: 3}, CellCoord{X: 3, Y: 4}, CellCoord{X: 3, Y: 5}, CellCoord{X: 4, Y: 3}, CellCoord{X: 4, Y: 4}, CellCoord{X: 4, Y: 5}, CellCoord{X: 5, Y: 3}, CellCoord{X: 5, Y: 4}, CellCoord{X: 5, Y: 5}},
	Block{CellCoord{X: 3, Y: 6}, CellCoord{X: 3, Y: 7}, CellCoord{X: 3, Y: 8}, CellCoord{X: 4, Y: 6}, CellCoord{X: 4, Y: 7}, CellCoord{X: 4, Y: 8}, CellCoord{X: 5, Y: 6}, CellCoord{X: 5, Y: 7}, CellCoord{X: 5, Y: 8}},
	Block{CellCoord{X: 6, Y: 0}, CellCoord{X: 6, Y: 1}, CellCoord{X: 6, Y: 2}, CellCoord{X: 7, Y: 0}, CellCoord{X: 7, Y: 1}, CellCoord{X: 7, Y: 2}, CellCoord{X: 8, Y: 0}, CellCoord{X: 8, Y: 1}, CellCoord{X: 8, Y: 2}},
	Block{CellCoord{X: 6, Y: 3}, CellCoord{X: 6, Y: 4}, CellCoord{X: 6, Y: 5}, CellCoord{X: 7, Y: 3}, CellCoord{X: 7, Y: 4}, CellCoord{X: 7, Y: 5}, CellCoord{X: 8, Y: 3}, CellCoord{X: 8, Y: 4}, CellCoord{X: 8, Y: 5}},
	Block{CellCoord{X: 6, Y: 6}, CellCoord{X: 6, Y: 7}, CellCoord{X: 6, Y: 8}, CellCoord{X: 7, Y: 6}, CellCoord{X: 7, Y: 7}, CellCoord{X: 7, Y: 8}, CellCoord{X: 8, Y: 6}, CellCoord{X: 8, Y: 7}, CellCoord{X: 8, Y: 8}},
}

var Cells []*Cell

type ThingType int

const (
	RowThing    ThingType = 0
	ColumnThing ThingType = iota
	BlockThing  ThingType = iota
)
