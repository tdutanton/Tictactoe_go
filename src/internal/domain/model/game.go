package model

import "github.com/google/uuid"

// Game - структура игры
type Game struct {
	ID    string
	Board Board
}

// NewGame - функция создания новой игры
// return *Game - указатель на экземпляр структуры Game
func NewGame() *Game {
	return &Game{
		ID:    uuid.NewString(),
		Board: Board{},
	}
}
