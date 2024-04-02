package git

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CloneGitRepo(url string) (*git.Repository, error) {
	r, err := git.PlainClone("/tmp/git-activity", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func CreateActiviyOnDayOfYear(r *git.Repository, dayOfYear time.Time, itensity int) error {
	w, err := r.Worktree()

	directory := "/tmo/git-activity"

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.
	fmt.Println("echo \"hello world!\" > example-git-file")
	filename := filepath.Join(directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0644)

	// Adds the new file to the staging area.
	fmt.Println("git add example-git-file")
	_, err = w.Add("example-git-file")

	// We can verify the current status of the worktree using the method Status.
	fmt.Println("git status --porcelain")
	status, err := w.Status()

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	fmt.Println("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	fmt.Println(err)

	// Prints the current HEAD to verify that all worked well.
	fmt.Println("git show -s")
	obj, err := r.CommitObject(commit)

	fmt.Println(obj)
	// create a commit on the given day of the year
	return nil
}

func PushGitRepo(r *git.Repository) error {
	// Push using default options
	err := r.Push(&git.PushOptions{})

	return err
}
