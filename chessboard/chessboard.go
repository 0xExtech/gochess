package chessboard

import . "gochess/piece"

type ChessBoard struct {
	Board   [65]int
	PiecesList [33]Piece
}

func NewChessBoard() ChessBoard {
    return ChessBoard{InitializeBoard(), InitializePiecesList()};
}

func InitializePiecesList() [33]Piece {
    var pieces  [33]Piece

    // Piece Null

    pieces[0] = PieceNull{0, true, false}

    // White Side

    pieces[1] = Rook{1, true, true}
    pieces[2] = Knight{2, true, true}
    pieces[3] = Bishop{3, true, true}
    pieces[4] = Queen{4, true, true}
    pieces[5] = King{5, true, true}
    pieces[6] = Bishop{6, true, true}
    pieces[7] = Knight{7, true, true}
    pieces[8] = Rook{8, true, true}

    for i := 9; i <= 16; i++ {
        pieces[i] = Pawn{i, true, true}
    }

    // Black Side

    pieces[17] = Rook{57, false, true}
    pieces[18] = Knight{58, false, true}
    pieces[19] = Bishop{59, false, true}
    pieces[20] = Queen{60, false, true}
    pieces[21] = King{61, false, true}
    pieces[22] = Bishop{62, false, true}
    pieces[23] = Knight{63, false, true}
    pieces[24] = Rook{64, false, true}

    for i, j := 25, 49; i <= 32; i, j = i + 1, j + 1 {
        pieces[i] = Pawn{j, true, true}
    }

    return pieces
}

func InitializeBoard() [65]int {
    var board   [65]int

    // White Pieces
    for i := 1; i <=8; i++ {
        board[i] = i
    }
    // White Pawns
    for i:= 9; i <= 16; i++ {
        board[i] = i
    }

    // Black Pieces
    for i, j := 57, 17; i <= 64; i, j = i + 1, j + 1 {
        board[i] = j
    }
    // Black Pawns
    for i, j := 49, 25; i <= 56; i, j = i + 1, j + 1 {
        board[i] = j
    }

    return board
}
