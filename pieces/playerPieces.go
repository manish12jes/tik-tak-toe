package pieces

type Pieces string

// type PieceX Pieces

// type PieceO Pieces

func NewPiece(val Pieces) *Pieces {
	piece := new(Pieces)
	*piece = val
	return piece
}
