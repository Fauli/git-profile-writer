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
	matrix := [][]int{
		//2    2    2    4	  5    6    7    8    9    20   22   22   22   24   25   26   27   28   29   20   22   22   22   24   25   26   27
		{0, 5, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 5, 5, 5, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{1, 5, 1, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 5, 5, 5, 5, 5, 1, 5, 5, 5, 5, 5, 1, 5, 1},
	}

	// print the matrix
	PrintMatrix(matrix)

	// initialize a date with the second of the year of the github activity year
	// TODO: Fauli: It should actially just be the first monday of the year
	date := time.Date(GITHUB_ACTIVTY_YEAR, time.January, 5, 0, 0, 0, 0, time.UTC)

	// Clone the git repository
	r, err := git.CloneGitRepo(GITHUB_REPO_URL)
	if err != nil {
		panic(err)
	}

	// to make the github activity matrix, we need to transpose the matrix
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			// create a github activity on the day of the year with the intensity of the matrix value
			err = git.CreateActiviyOnDayOfYear(r, date, int(matrix[j][i]))
			if err != nil {
				panic(err)
			}

			// increase the github activity date by one day
			date = date.AddDate(0, 0, 1)

		}
		fmt.Println()
	}

}

// TODO: Fauli:  This function should take a string and return a matrix that spells out the string
func StringToMatrix(s string) [][]int {
	return nil
}

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
