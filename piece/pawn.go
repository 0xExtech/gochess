package piece

import (
    "fmt"
)

type Pawn struct {
        Position int
        Color bool
        IsAlive bool
}

func (pawn Pawn) Move() {
        fmt.Println("Pawn moving")
}

func (pawn Pawn) Attack() {
}
