package matrixoperations

// RowSum - сумма значений в строке матрицы
func RowSum(rowNum int, matrix [3][3]int) int {
	result := 0
	for i := range 3 {
		result += matrix[rowNum][i]
	}
	return result
}

// RotateMatrix - поворот матрицы 3х3 на 90 градусов
func RotateMatrix(matrix [3][3]int) [3][3]int {
	var rotated [3][3]int
	for i := range 3 {
		for j := range 3 {
			rotated[i][j] = matrix[3-j-1][i]
		}
	}
	return rotated
}

// DiagSum - сумма значений по диагонали [0,0 - 2,2] матрицы 3х3
func DiagSum(matrix [3][3]int) int {
	result := 0
	for i := range 3 {
		result += matrix[i][i]
	}
	return result
}
