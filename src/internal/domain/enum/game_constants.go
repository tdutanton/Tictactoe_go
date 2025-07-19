package enum

const (
	// EMPTY - пустая ячейка
	EMPTY = 0
	// CROSS - крестик
	CROSS = 1
	// ZERO - нолик
	ZERO = -1

	// CrossWinSum - выигрышь крестиков (сумма ряда)
	CrossWinSum = CROSS * 3
	// ZeroWinSum - выигрышь ноликов (сумма ряда)
	ZeroWinSum = ZERO * 3
)
