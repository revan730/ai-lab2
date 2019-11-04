package main

type chessboardMoves interface {
	initialize(nSize int)
	squareIsFree(rowI int) bool
	setQueen(rowI int)
	removeQueen(rowI int)
	PlaceNextQueen()
}

type chessboard struct {
	columns      []bool // Store available column moves/attacks
	diagonalUp   []bool // Store available upward diagonal moves/attacks
	diagonalDown []bool // Store available downward diagonal moves/attacks
	columnJ      int    // Stores column to place the next queen in
}

func (board *chessboard) initialize() {
	for i := 0; i < cap(board.columns); i++ {
		board.columns[i] = true
	}
	for i := 0; i < cap(board.diagonalUp); i++ {
		board.diagonalUp[i] = true
	}
	copy(board.diagonalDown, board.diagonalUp)
}

func (board *chessboard) squareIsFree(rowI int) bool {
	return board.columns[rowI] &&
		board.diagonalUp[cap(board.columns)-1+board.columnJ-rowI] &&
		board.diagonalDown[board.columnJ+rowI]
}

func (board *chessboard) setQueen(rowI int) {
	board.columns[rowI] = false
	board.diagonalUp[cap(board.columns)-1+board.columnJ-rowI] = false
	board.diagonalDown[board.columnJ+rowI] = false
	board.columnJ++
}

func (board *chessboard) removeQueen(rowI int) {
	board.columnJ--
	board.diagonalDown[board.columnJ+rowI] = true
	board.diagonalUp[cap(board.columns)-1+board.columnJ-rowI] = true
	board.columns[rowI] = true
}

func main() {
	board := InitializeSolver(8)
	board.PlaceNextQueen()
}
