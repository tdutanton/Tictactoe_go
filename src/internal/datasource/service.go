package datasource

import (
	"errors"

	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

// GameRepositoryImpl — реализация GameRepository
type gameRepositoryImpl struct {
	store *GameStore
}

// NewGameRepository - создание нового сервиса репозитория
func NewGameRepository(store *GameStore) GameRepository {
	return &gameRepositoryImpl{
		store: store,
	}
}

// Save сохраняет model.Game в хранилище
func (r *gameRepositoryImpl) Save(game *model.Game) error {
	dsGame := MapToGameDS(game)
	r.store.Save(dsGame)
	return nil
}

// Find возвращает model.Game по ID
func (r *gameRepositoryImpl) Find(id string) (*model.Game, error) {
	dsGame, err := r.store.Get(id)
	if err != nil {
		return nil, errors.New("game not found")
	}
	return MapToDomainGame(dsGame), nil
}
