package main

import (
    "gochess/chessboard"
    "fmt"
)

func main() {
    var chess = chessboard.NewChessBoard()
    fmt.Printf("%v %T\n", chess.PiecesList[2], chess.PiecesList[2])
    fmt.Printf("%v %T\n", chess.PiecesList[5], chess.PiecesList[5])
    fmt.Printf("%v %T\n", chess.PiecesList[16], chess.PiecesList[16])
    fmt.Printf("%v %T\n", chess.PiecesList[23], chess.PiecesList[23])
    fmt.Printf("%v %T\n", chess.PiecesList[31], chess.PiecesList[31])
}
