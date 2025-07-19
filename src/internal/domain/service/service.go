package service

import (
	"errors"

	"github.com/tdutanton/Tictactoe_go/internal/datasource"
	"github.com/tdutanton/Tictactoe_go/internal/domain/enum"
	"github.com/tdutanton/Tictactoe_go/internal/domain/minimax"
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

type gameServiceImpl struct {
	repo datasource.GameRepository
}

// NewGameService - создание нового сервиса игры
func NewGameService(repo datasource.GameRepository) GameService {
	return &gameServiceImpl{repo: repo}
}

// NextMove - следующий ход
func (s *gameServiceImpl) NextMove(game *model.Game) (*model.Game, error) {
	board := game.Board
	bestMove := minimax.FindBestMove(board)
	if !board.IsEmptyCell(bestMove.X, bestMove.Y) {
		return nil, errors.New("cell is not empty")
	}
	board[bestMove.X][bestMove.Y] = enum.ZERO
	return &model.Game{
		ID:    game.ID,
		Board: board,
	}, nil
}

// Validate - проверка на изменение предыдущих ходов
func (s *gameServiceImpl) Validate(id string, newBoard model.Board) error {
	savedGame, err := s.repo.Find(id)
	if err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !savedGame.Board.IsEmptyCell(i, j) && savedGame.Board[i][j] != newBoard[i][j] {
				return errors.New("cannot modify non-empty cells")
			}
		}
	}
	diffCount := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if savedGame.Board[i][j] != newBoard[i][j] {
				diffCount++
				if diffCount > 1 {
					return errors.New("only one move is allowed per turn")
				}
			}
		}
	}
	if diffCount == 0 {
		return errors.New("no moves made")
	}
	return nil
}

// IsGameOver - проверка на конец игры
func (s *gameServiceImpl) IsGameOver(game *model.Game) bool {
	return game.Board.IsGameBoardOver()
}

// GetGame - геттер игры по id
func (s *gameServiceImpl) GetGame(id string) (*model.Game, error) {
	return s.repo.Find(id)
}

// SaveGame - сохранение игры
func (s *gameServiceImpl) SaveGame(game *model.Game) error {
	return s.repo.Save(game)
}
