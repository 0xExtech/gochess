package piece

import (
    "fmt"
)

type Bishop struct {
        Position int
        Color bool
        IsAlive bool
}

func (bishop Bishop) Move() {
        fmt.Println("Bishop moving")
}

func (bishop Bishop) Attack() {
}
