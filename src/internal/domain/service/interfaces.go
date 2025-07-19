package service

import (
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

// GameService - интерфейс для реализации сервисов игры
type GameService interface {
	NextMove(game *model.Game) (*model.Game, error)
	Validate(id string, newBoard model.Board) error
	IsGameOver(game *model.Game) bool
	GetGame(id string) (*model.Game, error)
	SaveGame(game *model.Game) error
}
