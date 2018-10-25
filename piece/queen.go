package piece

import (
    "fmt"
)

type Queen struct {
        Position int
        Color bool
        IsAlive bool
}

func (queen Queen) Move() {
        fmt.Println("Queen moving")
}

func (queen Queen) Attack() {
}
