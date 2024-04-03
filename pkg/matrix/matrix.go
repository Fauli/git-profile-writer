package matrix

import "fmt"

// NewMatrix creates a new matrix with the given number of rows and columns
func NewMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// PrintMatrix prints the matrix to the console
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("░")
			case 2:
				fmt.Print("▒")
			case 3:
				fmt.Print("▓")
			case 4:
				fmt.Print("█")
			case 5:
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
}

// TODO: Fauli:  This function should take a string and return a matrix that spells out the string
func StringToMatrix(s string) [][]int {
	return nil
}
