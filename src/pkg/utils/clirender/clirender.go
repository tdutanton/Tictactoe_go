package clirender

import (
	"fmt"

	"github.com/tdutanton/Tictactoe_go/internal/domain/enum"
	"github.com/tdutanton/Tictactoe_go/internal/domain/model"
)

const (
	// EMPTYSYMB - символ пустой ячейки для отрисовки
	EMPTYSYMB = ' '
	// CROSSSYMB - символ крестиков для отрисовки
	CROSSSYMB = 'X'
	// ZEROSYMB - символ ноликов для отрисовки
	ZEROSYMB = 'O'
)

// Symbols - map для сопоставления enum и символов
var Symbols = map[int]rune{
	enum.EMPTY: EMPTYSYMB,
	enum.CROSS: CROSSSYMB,
	enum.ZERO:  ZEROSYMB,
}

// RenderBoard - отрисовка доски
func RenderBoard(b model.Board) {
	for _, r := range b {
		for _, c := range r {
			fmt.Printf("%c", Symbols[c])
		}
		fmt.Println()
	}
}
