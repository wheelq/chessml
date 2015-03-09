package chess

import "fmt"

// PossibleMoves returns the possible moves of piece p from
// the position row,col.
func PossibleMoves(b Board, p Piece, row, col int) []ValidMove {
	switch p.Rank {
	case King:
		return KingMoves(b, row, col)
	case Queen:
		return QueenMoves(b, row, col)
	case Bishop:
		return BishopMoves(b, row, col)
	case Knight:
		return KnightMoves(b, row, col)
	case Rook:
		return RookMoves(b, row, col)
	case Pawn:
		return PawnMoves(b, row, col)
	default:
		panic(fmt.Sprintf("Invalid piece: bad Rank value %d", p.Rank))
	}
}

// RookMoves computes possible moves for a rook on board b at the row and col position.
func RookMoves(b Board, row, col int) (possible []ValidMove) {
	possible = append(possible, lineMove(b, row, col, -1, 0)...) // up
	possible = append(possible, lineMove(b, row, col, 1, 0)...)  // down
	possible = append(possible, lineMove(b, row, col, 0, -1)...) // left
	possible = append(possible, lineMove(b, row, col, 0, 1)...)  // right

	return
}

// BishopMoves computes possible moves for a bishop on board b at the row and col position.
func BishopMoves(b Board, row, col int) (possible []ValidMove) {
	possible = append(possible, lineMove(b, row, col, -1, -1)...) // up-left
	possible = append(possible, lineMove(b, row, col, -1, 1)...)  // up-right
	possible = append(possible, lineMove(b, row, col, 1, -1)...)  // down-left
	possible = append(possible, lineMove(b, row, col, 1, 1)...)   // down-right

	return
}

// QueenMoves computes possible moves for a queen on board b at the row and col position.
func QueenMoves(b Board, row, col int) (possible []ValidMove) {
	possible = append(possible, lineMove(b, row, col, -1, 0)...)  // up
	possible = append(possible, lineMove(b, row, col, 1, 0)...)   // down
	possible = append(possible, lineMove(b, row, col, 0, -1)...)  // left
	possible = append(possible, lineMove(b, row, col, 0, 1)...)   // right
	possible = append(possible, lineMove(b, row, col, -1, -1)...) // up-left
	possible = append(possible, lineMove(b, row, col, -1, 1)...)  // up-right
	possible = append(possible, lineMove(b, row, col, 1, -1)...)  // down-left
	possible = append(possible, lineMove(b, row, col, 1, 1)...)   // down-right

	return
}

// KnightMoves computes possible moves for a knight on board b at the row and col position.
func KnightMoves(b Board, row, col int) (possible []ValidMove) {
	possible = tryAndAppend(possible, b, row, col, -2, -1) // up-left
	possible = tryAndAppend(possible, b, row, col, -2, 1)  // up-right
	possible = tryAndAppend(possible, b, row, col, 2, -1)  // down-left
	possible = tryAndAppend(possible, b, row, col, 2, 1)   // down-right

	possible = tryAndAppend(possible, b, row, col, -1, -2) // left-up
	possible = tryAndAppend(possible, b, row, col, 1, -2)  // left-down
	possible = tryAndAppend(possible, b, row, col, -1, 2)  // right-up
	possible = tryAndAppend(possible, b, row, col, 1, 2)   // right-down

	return
}

// KingMoves computes possible moves for a king on board b at the row and col position.
func KingMoves(b Board, row, col int) (possible []ValidMove) {
	possible = tryAndAppend(possible, b, row, col, -1, 0) // up
	possible = tryAndAppend(possible, b, row, col, 1, 0)  // down
	possible = tryAndAppend(possible, b, row, col, 0, -1) // left
	possible = tryAndAppend(possible, b, row, col, 0, 1)  // right

	possible = tryAndAppend(possible, b, row, col, -1, -1) // up-left
	possible = tryAndAppend(possible, b, row, col, -1, 1)  // up-right
	possible = tryAndAppend(possible, b, row, col, 1, -1)  // down-left
	possible = tryAndAppend(possible, b, row, col, 1, 1)   // down-right

	return
}

// PawnMoves computes possible moves for a pawn on board b at the row and col position.
func PawnMoves(b Board, row, col int) (possible []ValidMove) {
	color := b.Spaces[row][col].Color

	if color == WhiteTeam {
		// move up (+)
		valid, _ := tryMove(b, color, row+1, col)

		if valid {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col}, Capture: false})
		}

		valid, capture := tryMove(b, color, row+1, col-1) // left capture

		if valid && capture {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col - 1}, Capture: true})
		}

		valid, capture = tryMove(b, color, row+1, col+1) // right capture

		if valid && capture {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col + 1}, Capture: true})
		}

		if row == 1 { // double move from starting row
			valid, capture = tryMove(b, color, row+2, col)

			if valid && !capture {
				possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 2, col}, Capture: true})
			}
		}

		if col+1 < Size && b.Spaces[row][col+1].EnPassantable {
			valid, capture = tryMove(b, color, row+1, col+1)

			if valid && !capture {
				possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col + 1}, EnPassant: true})
			}
		}

		if col-1 >= 0 && b.Spaces[row][col-1].EnPassantable {
			valid, capture = tryMove(b, color, row+1, col-1)

			if valid && !capture {
				possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col - 1}, EnPassant: true})
			}
		}

	} else {
		// move down (-)
		valid, _ := tryMove(b, color, row-1, col)

		if valid {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row - 1, col}, Capture: false})
		}

		valid, capture := tryMove(b, color, row-1, col-1) // left capture

		if valid && capture {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row - 1, col - 1}, Capture: true})
		}

		valid, capture = tryMove(b, color, row-1, col+1) // right capture

		if valid && capture {
			possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row - 1, col + 1}, Capture: true})
		}

		if row == 6 { // double move from starting row
			valid, capture = tryMove(b, color, row-2, col)

			if valid && !capture {
				possible = append(possible, ValidMove{
					From:    Coord{row, col},
					To:      Coord{row - 2, col},
					Capture: true,
				})
			}
		}

		if col+1 < Size && b.Spaces[row][col+1].EnPassantable {
			valid, capture = tryMove(b, color, row-1, col+1)

			if valid && !capture {
				possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col + 1}, EnPassant: true})
			}
		}

		if col-1 >= 0 && b.Spaces[row][col-1].EnPassantable {
			valid, capture = tryMove(b, color, row-1, col-1)

			if valid && !capture {
				possible = append(possible, ValidMove{From: Coord{row, col}, To: Coord{row + 1, col - 1}, EnPassant: true})
			}
		}

	}

	return
}

func tryMove(b Board, pieceColor Color, row, col int) (valid, capture bool) {
	if row < 0 || row >= Size || col < 0 || col >= Size {
		return false, false
	}

	target := b.Spaces[row][col]

	if target.Rank == Empty {
		// Valid move to empty square.
		return true, false
	} else if target.Color != pieceColor {
		// Enemy Piece captured.
		return true, true
	}

	return false, false
}

func lineMove(b Board, row, col, rowDiff, colDiff int) (possible []ValidMove) {
	color := b.Spaces[row][col].Color
	toRow, toCol := row, col

	for {
		toRow += rowDiff
		toCol += colDiff

		valid, capture := tryMove(b, color, toRow, toCol)

		if !valid {
			break
		} else if capture {
			possible = append(possible, ValidMove{
				From:    Coord{row, col},
				To:      Coord{toRow, toCol},
				Capture: capture,
			})
			break
		}

		possible = append(possible, ValidMove{
			From:    Coord{row, col},
			To:      Coord{toRow, toCol},
			Capture: capture,
		})
	}

	return
}

func tryAndAppend(vm []ValidMove, b Board, row, col, rowDiff, colDiff int) []ValidMove {
	color := b.Spaces[row][col].Color

	valid, capture := tryMove(b, color, row+rowDiff, col+colDiff)
	if valid {
		return append(vm, ValidMove{
			From:    Coord{row, col},
			To:      Coord{row + rowDiff, col + colDiff},
			Capture: capture,
		})
	}

	return vm
}

// ValidMove represents a possible move that has not necessarily been made.
type ValidMove struct {
	From, To  Coord
	Capture   bool
	EnPassant bool
}
