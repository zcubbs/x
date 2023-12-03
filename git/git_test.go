package git

import (
	"os"
	"testing"
)

func TestCloneRepository(t *testing.T) {
	const testRepoURL = "https://github.com/zcubbs/haproxy-test-repo-public"

	// Create a temporary directory for cloning
	tempDir, err := os.MkdirTemp("", "haproxy-test-repo-clone")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	err = CloneRepository(testRepoURL, tempDir, nil)
	if err != nil {
		t.Fatalf("CloneRepository failed: %s", err)
	}
}

func TestPullRepository(t *testing.T) {
	// Setup: Clone the repository first
	tempDir, err := os.MkdirTemp("", "haproxy-test-repo-pull")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	err = CloneRepository("https://github.com/zcubbs/haproxy-test-repo-public", tempDir, nil)
	if err != nil {
		t.Fatalf("Setup clone failed: %s", err)
	}

	// Test PullRepository
	_, err = PullRepository(tempDir, nil)
	if err != nil {
		t.Fatalf("PullRepository failed: %s", err)
	}
}

func TestGetLatestCommit(t *testing.T) {
	// Setup: Clone the repository first
	tempDir, err := os.MkdirTemp("", "haproxy-test-repo-commit")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	err = CloneRepository("https://github.com/zcubbs/haproxy-test-repo-public", tempDir, nil)
	if err != nil {
		t.Fatalf("Setup clone failed: %s", err)
	}

	commit, err := GetLatestCommit(tempDir)
	if err != nil {
		t.Fatalf("GetLatestCommit failed: %s", err)
	}

	if commit == "" {
		t.Error("Expected a commit hash, got empty string")
	}
}
