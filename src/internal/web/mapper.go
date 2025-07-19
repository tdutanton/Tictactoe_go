package web

import (
	"github.com/tdutanton/Tictactoe_go/internal/domain/enum"
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

// ToGameResponse - перевод из Domain в DTO
func ToGameResponse(game *model.Game) GameResponse {
	status := "in_progress"
	e, s := game.Board.GetGameBoardStatus()
	if e != enum.INPROGRESS {
		switch s {
		case enum.WINCROSS:
			status = "win_x"
		case enum.WINZERO:
			status = "win_o"
		case enum.TIE:
			status = "tie"
		}
	}

	var symbolBoard [3][3]string
	for i, row := range game.Board {
		for j, cell := range row {
			switch cell {
			case 1:
				symbolBoard[i][j] = "X"
			case -1:
				symbolBoard[i][j] = "0"
			default:
				symbolBoard[i][j] = " "
			}
		}
	}

	return GameResponse{
		ID:     game.ID,
		Board:  symbolBoard,
		Status: status,
	}
}

// ToDomainBoard - перевод игрового поля из DTO в Domain
func ToDomainBoard(byteBoard [3][3]string) model.Board {
	var intBoard model.Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch byteBoard[i][j] {
			case "X", "x":
				intBoard[i][j] = 1
			case "0":
				intBoard[i][j] = -1
			default:
				intBoard[i][j] = 0
			}
		}
	}
	return intBoard
}

// ToDomainGame - перевод из DTO в Domain
func ToDomainGame(req MoveRequest, gameID string) *model.Game {
	return &model.Game{
		ID:    gameID,
		Board: model.Board(ToDomainBoard(req.Board)),
	}
}
