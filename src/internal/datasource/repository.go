package datasource

import (
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

// GameRepository — интерфейс для работы с хранилищем игр
type GameRepository interface {
	Save(game *model.Game) error
	Find(id string) (*model.Game, error)
}
