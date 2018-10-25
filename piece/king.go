package piece

import (
    "fmt"
)

type King struct {
        Position int
        Color bool
        IsAlive bool
}

func (king King) Move() {
        fmt.Println("King moving")
}

func (king King) Attack() {
}
