package enum

// GameStatus статус игры
type GameStatus int

const (
	// INPROGRESS - игра в процессе
	INPROGRESS GameStatus = iota
	// GAMEOVER - конец игры на доске
	GAMEOVER
)

// GameOverStatus статус конца игры (кто-то выиграл или ничья)
type GameOverStatus int

const (
	// WINCROSS - выигрышь крестиков
	WINCROSS GameOverStatus = iota
	// WINZERO - выигрышь ноликов
	WINZERO
	// TIE - ничья
	TIE
)
