package model

// Board - игровое поле
type Board [3][3]int

// Cell - получить значение ячейки по "координатам"
func (b Board) Cell(x, y int) int {
	return b[x][y]
}

// Coord - структура для описания координаты доски
type Coord struct {
	X, Y int
}
