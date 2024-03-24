package board

import (
	"fmt"
	symb "tik-tak-toe/pieces"
)

type Board [][]symb.Pieces

func NewBoard(n int) Board {
	b := [][]symb.Pieces{}
	board := Board(b)
	board = board.createEmptyBoard(n)
	return board
}

func (b Board) AddPieceToBoad(piece symb.Pieces, x, y int) {
	b[x][y] = piece
}

func (b Board) ValidMove(x, y int) bool {
	n := len(b)
	if x < 0 || x >= n && y < 0 || y >= n {
		return false
	}
	if b[x][y] != " " {
		return false
	}
	return true
}

func (b Board) DisplayBoard() {
	n := len(b)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print("|")
			}
			fmt.Print(b[i][j])
		}
		fmt.Println()
	}
}

func (b Board) createEmptyBoard(n int) Board {
	for i := 0; i < n; i++ {
		tmp := make([]symb.Pieces, n)
		for j := 0; j < n; j++ {
			tmp[j] = " "
		}
		b = append(b, tmp)
	}
	return b
}
