package model

import (
	"errors"

	"github.com/tdutanton/Tictactoe_go/internal/domain/enum"
	matrixoperations "github.com/tdutanton/Tictactoe_go/pkg/utils/matrixOperations"
)

// IsFilled - заполнена ли доска полностью крестиками или ноликами
func (b Board) IsFilled() bool {
	for i, row := range b {
		for j := range row {
			if b.IsEmptyCell(i, j) {
				return false
			}
		}
	}
	return true
}

// IsEmptyCell - проверка на пустую ячейку
// return bool (пустая - true, непустая - false)
func (b Board) IsEmptyCell(x, y int) bool {
	return b[x][y] == enum.EMPTY
}

// ContainsValidValues - проверка на корректность игрового поля
// return bool (true - на поле допустимые символы, false - недопустимые)
func (b Board) ContainsValidValues() bool {
	for _, r := range b {
		for _, c := range r {
			if c != enum.EMPTY && c != enum.CROSS && c != enum.ZERO {
				return false
			}
		}
	}
	return true
}

// WhoWin - определение победителя по суммам значений строк, столбцов и диагоналей
func (b Board) WhoWin() (enum.GameOverStatus, error) {
	for rotation := 0; rotation < 2; rotation++ {
		for i := range len(b) {
			rs := matrixoperations.RowSum(i, b)
			switch rs {
			case enum.CrossWinSum:
				return enum.WINCROSS, nil
			case enum.ZeroWinSum:
				return enum.WINZERO, nil
			}
		}
		ds := matrixoperations.DiagSum(b)
		switch ds {
		case enum.CrossWinSum:
			return enum.WINCROSS, nil
		case enum.ZeroWinSum:
			return enum.WINZERO, nil
		}
		b = matrixoperations.RotateMatrix(b)
	}
	return 0, errors.New("there's no winners. Game must go on")
}

// IsBoardTie - проверка на ничью
func (b Board) IsBoardTie() bool {
	_, err := b.WhoWin()
	if b.IsFilled() && err != nil {
		return true
	}
	return false
}

// IsGameBoardOver - проверка на конец игры
func (b Board) IsGameBoardOver() bool {
	status, _ := b.GetGameBoardStatus()
	return status == enum.GAMEOVER
}

// GetGameBoardStatus - текущий статус игры
// return GameStatus, GameOverStatus
// Если есть выигравший - возвращает GameStatus = GAMEOVER, GameOverStatus = winner (WINCROSS или WINZERO)
// Если ничья - возвращает GameStatus = GAMEOVER, GameOverStatus = TIE
// Если оба варианта не подходят - возвращает INPROGRESS, 0
func (b Board) GetGameBoardStatus() (enum.GameStatus, enum.GameOverStatus) {
	winner, err := b.WhoWin()
	if err == nil {
		return enum.GAMEOVER, winner
	}
	if b.IsFilled() {
		return enum.GAMEOVER, enum.TIE
	}
	return enum.INPROGRESS, 0
}

// IsSame - сравнение двух досок
func (b Board) IsSame(bTwin Board) bool {
	for i, r := range b {
		for j, c := range r {
			if c != bTwin[i][j] {
				return false
			}
		}
	}
	return true
}
