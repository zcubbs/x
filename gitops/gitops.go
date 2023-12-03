package gitops

import (
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	xGit "github.com/zcubbs/x/git"
)

// Manager manages Git operations and file synchronization.
/* Example:
package main

import (
    "fmt"
    "log"
    "your-module-name/gitops"
)

func main() {
    repoURL := "https://github.com/example-user/example-repo.git"
    destination := "./example-repo"

    // Initialize GitOpsManager with credentials
    manager := gitops.NewGitOpsManager("your-username", "your-personal-access-token")

    // Clone a repository
    if err := manager.CloneRepository(repoURL, destination); err != nil {
        log.Fatalf("Clone failed: %s", err)
    }
    fmt.Println("Repository cloned successfully.")

    // Pull the repository
    changes, err := manager.PullRepository(destination)
    if err != nil {
        log.Fatalf("Pull failed: %s", err)
    }
    if changes {
        fmt.Println("Repository updated.")
    } else {
        fmt.Println("Repository already up to date.")
    }

    // Define file synchronization mappings
    mappings := []gitops.FileSyncMapping{
        {Source: "file1.txt", Destination: "/path/to/local/file1.txt"},
        // Add more mappings as needed
    }

    // Check for changes and synchronize files
    changes, err := manager.PullRepository(destination)
    if err != nil {
        log.Fatalf("Pull failed: %s", err)
    }
    if changes {
        fmt.Println("Changes detected, synchronizing files...")

        // Synchronize files
        if err := manager.SyncFiles(destination, mappings); err != nil {
            log.Fatalf("Sync failed: %s", err)
        }
        fmt.Println("Files synchronized successfully.")
    } else {
        fmt.Println("No changes detected.")
    }

    // Additional operations can be added here as needed
}

*/
type Manager struct {
	Auth *gitHttp.BasicAuth
}

// NewGitOpsManager creates a new instance of GitOpsManager with optional authentication.
func NewGitOpsManager(username, password string) *Manager {
	var auth *gitHttp.BasicAuth
	if username != "" && password != "" {
		auth = &gitHttp.BasicAuth{
			Username: username, // Can be anything except an empty string for token-based authentication
			Password: password,
		}
	}
	return &Manager{Auth: auth}
}

// CloneRepository clones a Git repository into a given directory.
func (m *Manager) CloneRepository(repoURL, destination string) error {
	return xGit.CloneRepository(repoURL, destination, m.Auth)
}

// PullRepository updates the local copy of a Git repository and returns true if there were changes.
func (m *Manager) PullRepository(repoPath string) (bool, error) {
	return xGit.PullRepository(repoPath, m.Auth)
}

// SyncFiles synchronizes files from the repository to the local file system based on the provided mappings.
func (m *Manager) SyncFiles(repoPath string, mappings []FileSyncMapping) error {
	return SyncFiles(repoPath, mappings)
}
