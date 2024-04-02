package main

import (
	"fmt"
	"time"

	"sbebe.ch/git-profile-writer/pkg/git"
)

const (
	// the width of the matrix in the github profile
	GITHUB_ACTIVTY_WIDTH = 51
	// the height of the matrix
	GITHUB_ACTIVTY_HEIGHT = 7
	// what year to write the text in
	GITHUB_ACTIVTY_YEAR = 2022
	// the URL of the github repository to clone
	GITHUB_REPO_URL = "https://github.com/Fauli/git-profile-writer-output.git"
)

func main() {

	// the matrix variable creates a 52x7 matrix
	// matrix := NewMatrix(27, 7)

	// initialize the matrix with only 0s and 2s as characters, so that it spells out "FAULI"
	matrix := [][]rune{
		//2    2    2    4	  5    6    7    8    9    20   22   22   22   24   25   26   27   28   29   20   22   22   22   24   25   26   27
		{'0', '2', '2', '2', '2', '2', '0', '2', '2', '2', '2', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'0', '2', '0', '0', '0', '0', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'0', '2', '0', '0', '0', '0', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'0', '2', '2', '2', '2', '0', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'0', '2', '0', '0', '0', '0', '0', '2', '2', '2', '2', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'0', '2', '0', '0', '0', '0', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '2', '0', '2', '0', '0', '0', '0', '0', '2', '0'},
		{'1', '2', '1', '1', '1', '1', '1', '2', '1', '1', '1', '2', '1', '2', '2', '2', '2', '2', '1', '2', '2', '2', '2', '2', '1', '2', '1'},
	}

	// print the matrix
	PrintMatrix(matrix)

	// initialize a date with the first of the year of the github activity year
	date := time.Date(GITHUB_ACTIVTY_YEAR, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Clone the git repository
	r, err := git.CloneGitRepo(GITHUB_REPO_URL)
	if err != nil {
		panic(err)
	}

	// to make the github activity matrix, we need to transpose the matrix
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			// print the date in the format "2006-01-02"
			fmt.Print(date.Format("2006-01-02: "))
			fmt.Print(string(matrix[j][i]))
			fmt.Print(", ")
			// increase the github activity date by one day
			date = date.AddDate(0, 0, 1)

			git.CreateActiviyOnDayOfYear(r, date, int(matrix[j][i]))
		}
		fmt.Println()
	}

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
			case '2':
				fmt.Print("█")
			case '1':
				fmt.Print("░")
			}
		}
		fmt.Println()
	}
}
