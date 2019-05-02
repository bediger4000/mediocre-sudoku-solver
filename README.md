# Mediocre Sudoku Solver

My 83-year-old mother does sudoku puzzles.
I asked her to show me how she solves them.
I decided to write my own solver to understand the methods.

## Building

    $ cd $GOPATH/src
    $ git clone https://github.com/bediger4000/mediocre-sudoku-solver.git ./sudoku
    $ cd sudoku
    $ go build sudoku
    $ go build cmd/cmp.go  # sudoku-board specific comparison
    $ ./runtests

The `sudoku` executable reads boards from stdin,
writes progress and any solution on stdout.

Partial output boards can be used as input,
by editing out all the non-board material.

The `runtests` script will try to solve all files in `tests/` directory as sudoku inputs.
It will note any incomplete solutions, or inputs that don't have a corresponding desired output.

### Other things it can do

It can definitely read ".sdk" format inputs

#### PostScript representations

    $ ./sudoku -p in11.ps < tests/in11
    $ ./sudoku -C -p in11possible.ps < tests/in11

The "-p filename" flag will produce a printable PostScript file
representing the input board.
With a "-C" flag, you get small digits in unsolved cells that show you
the possible digits that could fit in that cell.

#### Validate Sudoku boards

    $ ./sudoku -v < desired/15
    Complete Solution, valid
    $ ./sudoku -v < desired/00
    Incomplete Solution, valid

#### Turn on advanced methods of solution

* -H - perform "hidden subset" elimination
* -N - perform "naked subset" elimination
* -X - perform "XY wing" elimination

The `runtests` script turns on all advanced methods.

Using flags to turn on/off any methods can get you to a partial solution.
Using the "-f" flag to print only the final board,
you can create PostScript to understand why only a partial solution results.

    $ ./sudoku -N -f < tests/15 > 15n
    $ ./sudoku -C -p 15n.ps < 15n

Printing file `15n.ps` will show you the partial solution,
complete with all possible digits for unsolved cells.
