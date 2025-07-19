package web

// MoveRequest - ход игрока (запрос)
type MoveRequest struct {
	Board [3][3]string `json:"board" binding:"required"`
}

// GameResponse - ответ с состоянием игры
type GameResponse struct {
	ID     string       `json:"id"`
	Board  [3][3]string `json:"board"`
	Status string       `json:"status"` // "in_progress", "win_x", "win_o", "tie"
}

// ErrorResponse - ошибка
type ErrorResponse struct {
	Error string `json:"error"`
}
