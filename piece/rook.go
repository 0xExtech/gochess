package piece

import (
    "fmt"
)

type Rook struct {
	Position int
	Color    bool
	IsAlive  bool
}

func (rook Rook) Move() {
        fmt.Println("Rook moving")
}

func (rook Rook) Attack() {
}

func NewRook(pos int, col bool) Rook {
    return Rook{pos, col, true}
}
