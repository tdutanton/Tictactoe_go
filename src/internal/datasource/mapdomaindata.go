package datasource

import "github.com/tdutanton/Tictactoe_go/internal/domain/model"

// MapToGameDS преобразует model.Game в datasource.GameDS
func MapToGameDS(game *model.Game) *GameDS {
	return &GameDS{
		ID:    game.ID,
		Board: game.Board,
	}
}

// MapToDomainGame преобразует datasource.GameDS в model.Game
func MapToDomainGame(ds *GameDS) *model.Game {
	return &model.Game{
		ID:    ds.ID,
		Board: ds.Board,
	}
}
