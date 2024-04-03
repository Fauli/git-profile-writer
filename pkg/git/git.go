package git

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type GitUser struct {
	Name  string
	Email string
}

func CloneGitRepo(url string) (*git.Repository, error) {
	authMethod, err := ssh.NewSSHAgentAuth("git")
	if err != nil {
		return nil, err
	}

	r, err := git.PlainClone("/tmp/git-activity", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     authMethod,
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func CreateActiviyOnDayOfYear(r *git.Repository, user GitUser, dayOfYear time.Time, itensity int) error {
	fmt.Println("CALLED CreateActiviyOnDayOfYear", dayOfYear, itensity)

	// if itensity is 0, we don't need to do anything
	if itensity == 0 {
		return nil
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	directory := "/tmp/git-activity/"
	name := fmt.Sprintf("%s%d", "example-git-file", dayOfYear.Day())

	filename := filepath.Join(directory, name)

	for i := 0; i < itensity; i++ {
		fmt.Printf("Run %d of %d\n", i+1, itensity)

		fmt.Println("Creating file: ", filename)
		err = os.WriteFile(filename, []byte(fmt.Sprintf("somevalue%d", itensity)), 0644)
		if err != nil {

			return err
		}

		// Adds the new file to the staging area.
		fmt.Println("git add", name)
		_, err = w.Add(name)
		if err != nil {
			return err
		}

		// We can verify the current status of the worktree using the method Status.
		fmt.Println("git status --porcelain")
		status, err := w.Status()
		if err != nil {
			return err
		}

		fmt.Println(status)

		// Commits the current staging area to the repository, with the  file
		fmt.Println("git commit -m \"example go-git commit\"")
		commit, err := w.Commit("example go-git commit", &git.CommitOptions{
			Author: &object.Signature{
				When:  dayOfYear,
				Name:  user.Name,
				Email: user.Email,
			},
		})

		fmt.Println(err)

		// Prints the current HEAD to verify that all worked well.
		obj, err := r.CommitObject(commit)
		if err != nil {
			return err
		}

		fmt.Println(obj)
		// create a commit on the given day of the year
	}
	return nil
}

func PushGitRepo(r *git.Repository) error {
	// Push using default options
	err := r.Push(&git.PushOptions{})

	return err
}
