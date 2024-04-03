package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"sbebe.ch/git-profile-writer/pkg/git"
	"sbebe.ch/git-profile-writer/pkg/matrix"
	"sbebe.ch/git-profile-writer/pkg/utils"
)

const (
	// what year to write the text in
	GITHUB_ACTIVTY_YEAR = 2022
	// the URL of the github repository to clone
	GITHUB_REPO_URL = "git@github.com:Fauli/git-profile-writer-output.git"
)

func main() {
	// initialize the needed variables from the environment
	gitUser := utils.ReadGitUser()
	gitUrl := utils.GetEnv("GIT_REPO_URL", GITHUB_REPO_URL)
	activityYearEnv := utils.GetEnv("GITHUB_ACTIVTY_YEAR", fmt.Sprintf("%d", GITHUB_ACTIVTY_YEAR))
	// string to int
	activityYear, err := strconv.Atoi(activityYearEnv)
	if err != nil {
		panic(err)
	}

	// initialize the matrix with only 0s and 2s as characters, so that it spells out "FAULI"
	// in the future, there should be a function generating the matrix from a string
	text := [][]int{
		// FAULI
		{0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 5, 5, 5, 5, 5, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0},
		{2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2},
		{0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0},
		{0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 5, 5, 5, 5, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0},
		{2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 5, 0, 0, 0, 0, 0, 5, 5, 5, 5, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2},
		{0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 5, 0, 0, 0, 0, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 5, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 5, 5, 5, 5, 5, 1, 5, 5, 5, 5, 5, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	// print the matrix in the terminal
	matrix.PrintMatrix(text)

	// ask the user if it looks ok
	fmt.Println("Creating github activity as shown")
	ok := utils.AskForconfirmation("Do you want to continue?")
	if !ok {
		fmt.Println("Exiting")
		os.Exit(1)
	}

	// initialize a date with the first of the year of the github activity year
	date := time.Date(activityYear, time.January, 1, 0, 0, 0, 0, time.UTC)
	date = utils.CenterTheText(text, date)

	// Clone the git repository
	r, err := git.CloneGitRepo(gitUrl)
	if err != nil {
		panic(err)
	}

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
