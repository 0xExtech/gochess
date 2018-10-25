package piece

type PieceNull struct {
        Position int
        Color bool
        IsAlive bool
}

func (piecenull PieceNull) Move() {
}

func (piecenull PieceNull) Attack() {
}
