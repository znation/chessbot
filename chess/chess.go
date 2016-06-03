package chess

import (
  "time"
)

const (
  Empty = iota
  Pawn = iota
  Rook = iota
  Knight = iota
  Bishop = iota
  Queen = iota
  King = iota
)

type Board struct {
  Id      int64
  Created int64
  White [8][8]int8 // uses the constants above
  Black [8][8]int8 // uses the constants above
}

type Game struct {
  Id      int64
  Created int64
  White   int64
  Black   int64
  Board   int64
}

func newBoard() Board {
  return Board{
    Created: time.Now().UnixNano(),
    White: [8][8]int8{
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn},
      {Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook},
    },
    Black: [8][8]int8{
      {Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook},
      {Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn, Pawn},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
      {Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
    },
  }
}

func newGame(white int64, black int64, board int64) Game {
  return Game{
    Created: time.Now().UnixNano(),
    White: white,
    Black: black,
    Board: board,
  }
}
