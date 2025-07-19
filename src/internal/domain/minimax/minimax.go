package minimax

import (
	"math"

	"github.com/tdutanton/Tictactoe_go/internal/domain/enum"
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

const (
	// XSCORE - очки при выигрыше крестиков для алгоритма минимакс
	XSCORE = -10
	// OSCORE - очки при выигрыше ноликов для алгоритма минимакс
	OSCORE = 10
	// TIESCORE - очки при ничьей для алгоритма минимакс
	TIESCORE = 0
)

// AvailableMoves - определение списка возможных ходов на доске
func AvailableMoves(b model.Board) []model.Coord {
	var result []model.Coord
	for i, r := range b {
		for j, c := range r {
			if c == enum.EMPTY {
				result = append(result, model.Coord{X: i, Y: j})
			}
		}
	}
	return result
}

// Evaluate - подсчет очков в случае конца игры для алгоритма минимакс
func Evaluate(b model.Board) int {
	chck, result := b.GetGameBoardStatus()
	if chck != enum.INPROGRESS {
		switch result {
		case enum.WINCROSS:
			return XSCORE
		case enum.WINZERO:
			return OSCORE
		default:
			return TIESCORE
		}
	}
	return 0
}

// Minimax - алгоритм для выбора лучшего хода
func Minimax(b model.Board, isMax bool) int {
	if b.IsGameBoardOver() {
		return Evaluate(b)
	}
	if isMax {
		best := -math.MaxInt
		for _, move := range AvailableMoves(b) {
			b[move.X][move.Y] = enum.ZERO
			score := Minimax(b, false)
			b[move.X][move.Y] = enum.EMPTY
			best = max(best, score)
		}
		return best
	}
	best := math.MaxInt
	for _, move := range AvailableMoves(b) {
		b[move.X][move.Y] = enum.CROSS
		score := Minimax(b, true)
		b[move.X][move.Y] = enum.EMPTY
		best = min(best, score)
	}
	return best
}

// FindBestMove - определение координат поля для лучшего хода противника
func FindBestMove(b model.Board) model.Coord {
	bestScore := -math.MaxInt
	var bestMove model.Coord
	moves := AvailableMoves(b)

	for _, move := range moves {
		b[move.X][move.Y] = enum.ZERO
		score := Minimax(b, false)
		b[move.X][move.Y] = enum.EMPTY

		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return bestMove
}
