package web

import (
	"net/http"

	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
	"github.com/tdutanton/Tictactoe_go/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter - настройка gin роутера
func SetupRouter(gameService service.GameService) *gin.Engine {
	router := gin.Default()
	gameHandler := NewGameHandler(gameService)

	api := router.Group("/api")
	{
		api.POST("/game/:id", gameHandler.MakeMove)
		api.POST("/game", func(c *gin.Context) {
			newGame := model.NewGame()
			if err := gameService.SaveGame(newGame); err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
				return
			}
			c.JSON(http.StatusCreated, ToGameResponse(newGame))
		})
	}

	return router
}
