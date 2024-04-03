package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"sbebe.ch/git-profile-writer/pkg/git"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	fmt.Printf("Using fallback for %s: %s\n", key, fallback)
	return fallback
}

func AskForconfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// skipToFirstSundayOfTheYear skips the date to the first Sunday of the year
func SkipToFirstSundayOfTheYear(date time.Time) time.Time {
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}

	fmt.Printf("Skipping to first Sunday of the year: %s\n", date)
	return date
}

// centerTheText skips the first few weeks, because the github activity has 52 weeks
// however, the matrix text with len(matrix[0]) might be shorter than 52 weeks.
// The text should be centered, so half the weeks should be skipped at the beginning
func CenterTheText(text [][]int, date time.Time) time.Time {

	// as the first line is sundays, skip to the first sunday of the year
	date = SkipToFirstSundayOfTheYear(date)

	// skip the first few weeks
	// the number of weeks to skip is the difference between the number of weeks in the matrix and the number of weeks in the year
	weeksToSkip := (52 - len(text[0])) / 2
	date = date.AddDate(0, 0, weeksToSkip*7)

	// skip days until we are at a Sunday again
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}
	fmt.Printf("Skipping to start for centered Sunday of the year: %s\n", date)

	return date
}

func ReadGitUser() git.GitUser {
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
