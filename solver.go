package main

import "fmt"

type solverBoard struct {
	chessboard       
	queens     []int
}

func InitializeSolver(nSize int) *solverBoard {
	diagonalLength := 2*nSize - 1
	board := &solverBoard{
		chessboard: chessboard{
			columns:      make([]bool, nSize, nSize),
			diagonalUp:   make([]bool, diagonalLength, diagonalLength),
			diagonalDown: make([]bool, diagonalLength, diagonalLength),
			columnJ:      0,
		},
		queens: make([]int, nSize, nSize),
	}
	board.initialize()

	return board
}

func (board *solverBoard) PlaceNextQueen() {
	for rowI := 0; rowI < cap(board.columns); rowI++ {
		if board.squareIsFree(rowI) {
			board.queens[board.columnJ] = rowI + 1
			board.setQueen(rowI)
			if board.columnJ == cap(board.columns) {
				// Chess board is full
				for row := 0; row < cap(board.columns); row++ {
					fmt.Printf("%d ", board.queens[row])
				}
				fmt.Printf("\n")
			} else {
				board.PlaceNextQueen()
			}
			board.removeQueen(rowI)
		}
	}
}
