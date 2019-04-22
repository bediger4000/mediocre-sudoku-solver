# Mediocre Sudoku Solver

My 83-year-old mom does sudoku puzzles.
I asked her to show me how she solves them.
I decided to write my own solver to understand the methods.
I have departed from her solution methods.

## Building

   $ cd $GOPATH/src
   $ git clone https://github.com/bediger4000/mediocre-sudoku-solver.git
   $ cd sudoku
   $ go build sudoku
   $ go build cmd/cmp.go
   $ ./sudoku < tests/in17

The `sudoku` executable reads boards from stdin,
writes progress and any solution on stdout.

Partial output boards can be used as input,
by editing out all the non-board material.

### Other things it can do

   $ ./sudoku -C -p in11possible.ps < tests/in11

The "-p filename" flag will produce a printable PostScript file
representing the input board.
With a "-C" flag, you get small digits in unsolved cells that show you
the possible digits that could fit in that cell.

