package piece

import (
    "fmt"
)

type Knight struct {
        Position int
        Color bool
        IsAlive bool
}

func (knight Knight) Move() {
            fmt.Println("Knight moving")
}

func (knight Knight) Attack() {
}
