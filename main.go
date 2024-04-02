package main

import "fmt"

const (
	// the width of the matrix in the github profile
	GITHUB_ACTIVTY_WIDTH = 27
	// the height of the matrix
	GITHUB_ACTIVTY_HEIGHT = 7
)

func main() {

	// the matrix variable creates a 51x7 matrix
	// matrix := NewMatrix(27, 7)

	// initialize the matrix with only 0s and 1s as characters, so that it spells out "FAULI"
	matrix := [][]rune{
		//1    2    3    4	  5    6    7    8    9    10   11   12   13   14   15   16   17   18   19   20   21   22   23   24   25   26   27
		{'0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '1', '1', '1', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '1', '0', '1', '0', '0', '0', '0', '0', '1', '0'},
		{'0', '1', '0', '0', '0', '0', '0', '1', '0', '0', '0', '1', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0', '1', '0'},
	}

	// print the matrix
	PrintMatrix(matrix)

}

// TODO: Fauli:  This function should take a string and return a matrix that spells out the string
func StringToMatrix(s string) [][]rune {
	return nil
}

// NewMatrix creates a new matrix with the given number of rows and columns
func NewMatrix(rows, cols int) [][]rune {
	matrix := make([][]rune, rows)
	for i := range matrix {
		matrix[i] = make([]rune, cols)
	}
	return matrix
}

// PrintMatrix prints the matrix to the console
func PrintMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, cell := range row {
			switch cell {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
}
