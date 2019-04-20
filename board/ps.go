package board

import (
	"fmt"
	"io"
)

var digitOffsets [10]CellCoord = [10]CellCoord{
	CellCoord{0, 0},
	CellCoord{-17, 17},
	CellCoord{0, 17},
	CellCoord{17, 17},
	CellCoord{-17, 0},
	CellCoord{0, 0},
	CellCoord{17, 0},
	CellCoord{-17, -17},
	CellCoord{0, -17},
	CellCoord{17, -17},
}

func (bd *Board) EmitPostScript(out io.Writer, printPossibles bool) {
	out.Write([]byte(PSHeader))
	/*
		94 94 moveto
		(8) show
		144 94 moveto
		(9) show
		94 144 moveto
		(1) show
	*/
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			xoffset := 50*col + 94
			yoffset := 494 - 50*row
			if bd[row][col].Solved {
				out.Write([]byte(fmt.Sprintf("%d %d moveto\n(%d) show\n", xoffset, yoffset, bd[row][col].Value)))
			} else if printPossibles {
				out.Write([]byte("/Times-Roman findfont\n9 scalefont\nsetfont\n"))
				for _, digit := range bd[row][col].Possible {
					if digit == 0 {
						continue
					}
					digitOffset := digitOffsets[digit]
					out.Write([]byte(fmt.Sprintf("%d %d moveto\n(%d) show\n", xoffset+digitOffset.X, yoffset+digitOffset.Y, digit)))
				}
				out.Write([]byte("/Times-Roman findfont\n22 scalefont\nsetfont\n"))
			}
		}
	}

	out.Write([]byte("showpage"))
}

var PSHeader string = `%!PS

newpath
2 setlinewidth
 75 75 moveto
 450 0   rlineto
 0   450 rlineto
-450 0   rlineto
 0  -450 rlineto
stroke

newpath
1 setlinewidth
125 75 moveto
0  450 rlineto
50   0 rmoveto
0 -450 rlineto

47   0 rmoveto
0  450 rlineto
3    0 rmoveto
0 -450 rlineto

47   0 rmoveto
0  450 rlineto
50   0 rmoveto
0 -450 rlineto

47 0 rmoveto
0  450 rlineto
3    0 rmoveto
0 -450 rlineto

47   0 rmoveto
0 450 rlineto
50   0 rmoveto
0  -450 rlineto
stroke

newpath
1 setlinewidth
75 125 moveto
450  0 rlineto

0   50 rmoveto
-450 0 rlineto

0   47 rmoveto
450  0 rlineto
0    3 rmoveto
-450 0 rlineto

0   50 rmoveto
450 0 rlineto

0   50 rmoveto
-450  0 rlineto

0   47 rmoveto
450 0 rlineto
0    3 rmoveto
-450 0 rlineto

0   50 rmoveto
450  0 rlineto
0   50 rmoveto
-450 0 rlineto
0   50 rmoveto
450  0 rlineto
stroke

/Times-Roman findfont
22 scalefont
setfont
newpath
`
