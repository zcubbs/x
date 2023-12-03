// Package git provides utilities to interact with git repositories.
//
// This package is primarily designed to clone and pull updates from
// git repositories, specifically with support for optional credentials
// in the form of username and password from environment variables.
//
// Author: zakaria.elbouwab
// zcubbs https://github.com/zcubbs
package git

import (
	"errors"
	"fmt"
	git "github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/zcubbs/x/bash"
)

// CloneRepository clones a Git repository into a given directory with optional authentication.
func CloneRepository(repoURL, destination string, auth *gitHttp.BasicAuth) error {
	cloneOptions := &git.CloneOptions{
		URL:  repoURL,
		Auth: auth, // This can be nil, which is handled by go-git
	}

	_, err := git.PlainClone(destination, false, cloneOptions)
	if err != nil {
		return fmt.Errorf("error cloning repository: %w", err)
	}
	return nil
}

// PullRepository updates the local copy of a Git repository with optional authentication and returns true if there were changes.
func PullRepository(repoPath string, auth *gitHttp.BasicAuth) (bool, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return false, fmt.Errorf("error opening repository: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return false, fmt.Errorf("error getting worktree: %w", err)
	}

	pullOptions := &git.PullOptions{
		RemoteName: "origin",
		Auth:       auth,
	}

	err = w.Pull(pullOptions)
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return false, fmt.Errorf("error pulling repository: %w", err)
	}

	return !errors.Is(err, git.NoErrAlreadyUpToDate), nil
}

func GetLatestCommit(gitRepoPath string) (string, error) {
	commit, err := bash.ExecuteCmdWithOutput("git", "-C", gitRepoPath, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}

	return commit, nil
}
