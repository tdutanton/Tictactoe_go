package datasource

import (
	"errors"
	"sync"
)

// GameStore — потокобезопасное хранилище игр
type GameStore struct {
	data *sync.Map
}

// NewGameStore - создание нового хранилища
func NewGameStore() *GameStore {
	return &GameStore{
		data: &sync.Map{},
	}
}

// Save сохраняет игру в хранилище
func (s *GameStore) Save(game *GameDS) {
	s.data.Store(game.ID, game)
}

// Get возвращает игру по ID
func (s *GameStore) Get(id string) (*GameDS, error) {
	v, ok := s.data.Load(id)
	if !ok {
		return nil, errors.New("get game error")
	}
	return v.(*GameDS), nil
}
