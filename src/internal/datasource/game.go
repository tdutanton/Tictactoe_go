package datasource

// GameDS — модель текущей игры в datasource-слое
type GameDS struct {
	// ID - uuid игры
	ID string `json:"id"`
	// Board - игровое поле (0 - пусто, 1 - крести, -1 - нолик)
	Board [3][3]int `json:"board"`
}
