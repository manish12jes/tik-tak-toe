package player

import (
	"tik-tak-toe/board"
	symb "tik-tak-toe/pieces"
)

type Player struct {
	Name  string
	Piece symb.Pieces
}

func NewPlayer(name string, piece symb.Pieces) *Player {
	p := &Player{}
	p.Name = name
	p.Piece = piece
	return p
}
func (p Player) IsWinner(b board.Board, piece symb.Pieces) bool {
	var (
		rowMatch, colMatch, diaMatch1, diaMatch2 bool
		n                                        = len(b)
	)

	for i := 0; i < 3; i++ {
		rowMatch = true
		for j := 0; j < 3; j++ {
			if b[i][j] != piece {
				rowMatch = false
			}
		}
		if rowMatch {
			break
		}
	}
	for i := 0; i < 3; i++ {
		colMatch = true
		for j := 0; j < 3; j++ {
			if b[j][i] != piece {
				colMatch = false
			}
		}
		if colMatch {
			break
		}
	}
	diaMatch1, diaMatch2 = true, true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				if b[i][j] != piece {
					diaMatch1 = false
				}
			} else if i == n-j-1 && j == n-i-1 {
				if b[i][j] != piece {
					diaMatch2 = false
				}
			}
		}
	}
	return rowMatch || colMatch || diaMatch1 || diaMatch2
}
