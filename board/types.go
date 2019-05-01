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

type House [9]CellCoord

var Blocks [9]House = [9]House{
	House{CellCoord{X: 0, Y: 0}, CellCoord{X: 0, Y: 1}, CellCoord{X: 0, Y: 2}, CellCoord{X: 1, Y: 0}, CellCoord{X: 1, Y: 1}, CellCoord{X: 1, Y: 2}, CellCoord{X: 2, Y: 0}, CellCoord{X: 2, Y: 1}, CellCoord{X: 2, Y: 2}},
	House{CellCoord{X: 0, Y: 3}, CellCoord{X: 0, Y: 4}, CellCoord{X: 0, Y: 5}, CellCoord{X: 1, Y: 3}, CellCoord{X: 1, Y: 4}, CellCoord{X: 1, Y: 5}, CellCoord{X: 2, Y: 3}, CellCoord{X: 2, Y: 4}, CellCoord{X: 2, Y: 5}},
	House{CellCoord{X: 0, Y: 6}, CellCoord{X: 0, Y: 7}, CellCoord{X: 0, Y: 8}, CellCoord{X: 1, Y: 6}, CellCoord{X: 1, Y: 7}, CellCoord{X: 1, Y: 8}, CellCoord{X: 2, Y: 6}, CellCoord{X: 2, Y: 7}, CellCoord{X: 2, Y: 8}},
	House{CellCoord{X: 3, Y: 0}, CellCoord{X: 3, Y: 1}, CellCoord{X: 3, Y: 2}, CellCoord{X: 4, Y: 0}, CellCoord{X: 4, Y: 1}, CellCoord{X: 4, Y: 2}, CellCoord{X: 5, Y: 0}, CellCoord{X: 5, Y: 1}, CellCoord{X: 5, Y: 2}},
	House{CellCoord{X: 3, Y: 3}, CellCoord{X: 3, Y: 4}, CellCoord{X: 3, Y: 5}, CellCoord{X: 4, Y: 3}, CellCoord{X: 4, Y: 4}, CellCoord{X: 4, Y: 5}, CellCoord{X: 5, Y: 3}, CellCoord{X: 5, Y: 4}, CellCoord{X: 5, Y: 5}},
	House{CellCoord{X: 3, Y: 6}, CellCoord{X: 3, Y: 7}, CellCoord{X: 3, Y: 8}, CellCoord{X: 4, Y: 6}, CellCoord{X: 4, Y: 7}, CellCoord{X: 4, Y: 8}, CellCoord{X: 5, Y: 6}, CellCoord{X: 5, Y: 7}, CellCoord{X: 5, Y: 8}},
	House{CellCoord{X: 6, Y: 0}, CellCoord{X: 6, Y: 1}, CellCoord{X: 6, Y: 2}, CellCoord{X: 7, Y: 0}, CellCoord{X: 7, Y: 1}, CellCoord{X: 7, Y: 2}, CellCoord{X: 8, Y: 0}, CellCoord{X: 8, Y: 1}, CellCoord{X: 8, Y: 2}},
	House{CellCoord{X: 6, Y: 3}, CellCoord{X: 6, Y: 4}, CellCoord{X: 6, Y: 5}, CellCoord{X: 7, Y: 3}, CellCoord{X: 7, Y: 4}, CellCoord{X: 7, Y: 5}, CellCoord{X: 8, Y: 3}, CellCoord{X: 8, Y: 4}, CellCoord{X: 8, Y: 5}},
	House{CellCoord{X: 6, Y: 6}, CellCoord{X: 6, Y: 7}, CellCoord{X: 6, Y: 8}, CellCoord{X: 7, Y: 6}, CellCoord{X: 7, Y: 7}, CellCoord{X: 7, Y: 8}, CellCoord{X: 8, Y: 6}, CellCoord{X: 8, Y: 7}, CellCoord{X: 8, Y: 8}},
}

var Columns [9]House = [9]House{
	House{CellCoord{0, 0}, CellCoord{1, 0}, CellCoord{2, 0}, CellCoord{3, 0}, CellCoord{4, 0}, CellCoord{5, 0}, CellCoord{6, 0}, CellCoord{7, 0}, CellCoord{8, 0}},
	House{CellCoord{0, 1}, CellCoord{1, 1}, CellCoord{2, 1}, CellCoord{3, 1}, CellCoord{4, 1}, CellCoord{5, 1}, CellCoord{6, 1}, CellCoord{7, 1}, CellCoord{8, 1}},
	House{CellCoord{0, 2}, CellCoord{1, 2}, CellCoord{2, 2}, CellCoord{3, 2}, CellCoord{4, 2}, CellCoord{5, 2}, CellCoord{6, 2}, CellCoord{7, 2}, CellCoord{8, 2}},
	House{CellCoord{0, 3}, CellCoord{1, 3}, CellCoord{2, 3}, CellCoord{3, 3}, CellCoord{4, 3}, CellCoord{5, 3}, CellCoord{6, 3}, CellCoord{7, 3}, CellCoord{8, 3}},
	House{CellCoord{0, 4}, CellCoord{1, 4}, CellCoord{2, 4}, CellCoord{3, 4}, CellCoord{4, 4}, CellCoord{5, 4}, CellCoord{6, 4}, CellCoord{7, 4}, CellCoord{8, 4}},
	House{CellCoord{0, 5}, CellCoord{1, 5}, CellCoord{2, 5}, CellCoord{3, 5}, CellCoord{4, 5}, CellCoord{5, 5}, CellCoord{6, 5}, CellCoord{7, 5}, CellCoord{8, 5}},
	House{CellCoord{0, 6}, CellCoord{1, 6}, CellCoord{2, 6}, CellCoord{3, 6}, CellCoord{4, 6}, CellCoord{5, 6}, CellCoord{6, 6}, CellCoord{7, 6}, CellCoord{8, 6}},
	House{CellCoord{0, 7}, CellCoord{1, 7}, CellCoord{2, 7}, CellCoord{3, 7}, CellCoord{4, 7}, CellCoord{5, 7}, CellCoord{6, 7}, CellCoord{7, 7}, CellCoord{8, 7}},
	House{CellCoord{0, 8}, CellCoord{1, 8}, CellCoord{2, 8}, CellCoord{3, 8}, CellCoord{4, 8}, CellCoord{5, 8}, CellCoord{6, 8}, CellCoord{7, 8}, CellCoord{8, 8}},
}

var Rows [9]House = [9]House{
	House{CellCoord{0, 0}, CellCoord{0, 1}, CellCoord{0, 2}, CellCoord{0, 3}, CellCoord{0, 4}, CellCoord{0, 5}, CellCoord{0, 6}, CellCoord{0, 7}, CellCoord{0, 8}},
	House{CellCoord{1, 0}, CellCoord{1, 1}, CellCoord{1, 2}, CellCoord{1, 3}, CellCoord{1, 4}, CellCoord{1, 5}, CellCoord{1, 6}, CellCoord{1, 7}, CellCoord{1, 8}},
	House{CellCoord{2, 0}, CellCoord{2, 1}, CellCoord{2, 2}, CellCoord{2, 3}, CellCoord{2, 4}, CellCoord{2, 5}, CellCoord{2, 6}, CellCoord{2, 7}, CellCoord{2, 8}},
	House{CellCoord{3, 0}, CellCoord{3, 1}, CellCoord{3, 2}, CellCoord{3, 3}, CellCoord{3, 4}, CellCoord{3, 5}, CellCoord{3, 6}, CellCoord{3, 7}, CellCoord{3, 8}},
	House{CellCoord{4, 0}, CellCoord{4, 1}, CellCoord{4, 2}, CellCoord{4, 3}, CellCoord{4, 4}, CellCoord{4, 5}, CellCoord{4, 6}, CellCoord{4, 7}, CellCoord{4, 8}},
	House{CellCoord{5, 0}, CellCoord{5, 1}, CellCoord{5, 2}, CellCoord{5, 3}, CellCoord{5, 4}, CellCoord{5, 5}, CellCoord{5, 6}, CellCoord{5, 7}, CellCoord{5, 8}},
	House{CellCoord{6, 0}, CellCoord{6, 1}, CellCoord{6, 2}, CellCoord{6, 3}, CellCoord{6, 4}, CellCoord{6, 5}, CellCoord{6, 6}, CellCoord{6, 7}, CellCoord{6, 8}},
	House{CellCoord{7, 0}, CellCoord{7, 1}, CellCoord{7, 2}, CellCoord{7, 3}, CellCoord{7, 4}, CellCoord{7, 5}, CellCoord{7, 6}, CellCoord{7, 7}, CellCoord{7, 8}},
	House{CellCoord{8, 0}, CellCoord{8, 1}, CellCoord{8, 2}, CellCoord{8, 3}, CellCoord{8, 4}, CellCoord{8, 5}, CellCoord{8, 6}, CellCoord{8, 7}, CellCoord{8, 8}},
}

var Cells []*Cell

type ThingType int

const (
	RowThing    ThingType = 0
	ColumnThing ThingType = iota
	BlockThing  ThingType = iota
)
