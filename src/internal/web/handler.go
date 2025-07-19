package web

import (
	"net/http"

	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
	"github.com/tdutanton/Tictactoe_go/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// GameHandler - обработчик web-слоя игры
type GameHandler struct {
	gameService service.GameService
}

// NewGameHandler - новый хендлер
func NewGameHandler(gameService service.GameService) *GameHandler {
	return &GameHandler{gameService: gameService}
}

// MakeMove - команда для хода на доске
func (h *GameHandler) MakeMove(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "game ID is required"})
		return
	}
	var moveReq MoveRequest
	if err := c.ShouldBindJSON(&moveReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	// Валидация хода
	if err := h.gameService.Validate(gameID, ToDomainBoard(moveReq.Board)); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	// Получение текущей игры
	_, err := h.gameService.GetGame(gameID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	// Обновляется игра с ходом игрока
	updatedGame := &model.Game{
		ID:    gameID,
		Board: ToDomainBoard(moveReq.Board),
	}
	// Если игра не закончена, делается ход PC
	if !h.gameService.IsGameOver(updatedGame) {
		updatedGame, err = h.gameService.NextMove(updatedGame)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
	}
	// Сохранение обновленной игры
	if err := h.gameService.SaveGame(updatedGame); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToGameResponse(updatedGame))
}
