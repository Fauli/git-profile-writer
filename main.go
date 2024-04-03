package main

import (
	"fmt"
	"os"
	"time"

	"sbebe.ch/git-profile-writer/pkg/git"
	"sbebe.ch/git-profile-writer/pkg/matrix"
)

const (
	// the width of the matrix in the github profile
	GITHUB_ACTIVTY_WIDTH = 52
	// the height of the matrix
	GITHUB_ACTIVTY_HEIGHT = 7
	// what year to write the text in
	GITHUB_ACTIVTY_YEAR = 2022
	// the URL of the github repository to clone
	GITHUB_REPO_URL = "git@github.com:Fauli/git-profile-writer-output.git"
)

func main() {

	gitUser := readGitUser()

	// the matrix variable creates a 52x7 matrix
	// matrix := NewMatrix(27, 7)

	// initialize the matrix with only 0s and 2s as characters, so that it spells out "FAULI"
	text := [][]int{
		// FAULI
		{0, 5, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 5, 5, 5, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0},
		{1, 5, 1, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 5, 5, 5, 5, 5, 1, 5, 5, 5, 5, 5, 1, 5, 1},
	}

	// print the matrix
	matrix.PrintMatrix(text)

	// initialize a date with the second of the year of the github activity year
	// TODO: Fauli: It should actially just be the first Sunday of the year
	date := time.Date(GITHUB_ACTIVTY_YEAR, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Clone the git repository
	r, err := git.CloneGitRepo(GITHUB_REPO_URL)
	if err != nil {
		panic(err)
	}

	date = centerTheText(text, date)

	// to make the github activity matrix, we need to transpose the matrix
	for i := 0; i < len(text[0]); i++ {
		for j := 0; j < len(text); j++ {
			// create a github activity on the day of the year with the intensity of the matrix value
			err = git.CreateActiviyOnDayOfYear(r, gitUser, date, int(text[j][i]))
			if err != nil {
				panic(err)
			}

			// increase the github activity date by one day
			date = date.AddDate(0, 0, 1)

		}
		fmt.Println()
	}

	// push the changes to the repository
	err = git.PushGitRepo(r)
	if err != nil {
		panic(err)
	}

}

// skipToFirstSundayOfTheYear skips the date to the first Sunday of the year
func skipToFirstSundayOfTheYear(date time.Time) time.Time {
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}

	fmt.Printf("Skipping to first Sunday of the year: %s\n", date)
	return date
}

// centerTheText skips the first few weeks, because the github activity has 52 weeks
// however, the matrix text with len(matrix[0]) might be shorter than 52 weeks.
// The text should be centered, so half the weeks should be skipped at the beginning
func centerTheText(text [][]int, date time.Time) time.Time {

	// as the first line is sundays, skip to the first sunday of the year
	date = skipToFirstSundayOfTheYear(date)

	// skip the first few weeks
	// the number of weeks to skip is the difference between the number of weeks in the matrix and the number of weeks in the year
	weeksToSkip := (GITHUB_ACTIVTY_WIDTH - len(text[0])) / 2
	date = date.AddDate(0, 0, weeksToSkip*7)

	// skip days until we are at a Sunday again
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}
	fmt.Printf("Skipping to start for centered Sunday of the year: %s\n", date)

	return date
}

func readGitUser() git.GitUser {
	// Read GIT_USER and GIT_EMAIL from the environment
	// if they are not set, exit and print an error message
	user := os.Getenv("GIT_USER")
	if user == "" {
		fmt.Println("GIT_USER environment variable not set")
		os.Exit(1)
	}

	email := os.Getenv("GIT_EMAIL")
	if email == "" {
		fmt.Println("GIT_EMAIL environment variable not set")
		os.Exit(1)
	}

	return git.GitUser{
		Name:  user,
		Email: email,
	}
}
