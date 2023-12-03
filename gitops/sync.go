package gitops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileSyncMapping defines a mapping of a file from its source location in the repository to a destination path in the local file system.
type FileSyncMapping struct {
	Source      string
	Destination string
}

// SyncFiles synchronizes files from the repository to the local file system based on the provided mappings.
func SyncFiles(repoPath string, mappings []FileSyncMapping) error {
	for _, mapping := range mappings {
		srcPath := filepath.Join(repoPath, mapping.Source)
		destPath := filepath.Join(mapping.Destination)

		// Create the destination directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
			return fmt.Errorf("error creating destination directory: %w", err)
		}

		// Copy the file
		if err := copyFile(srcPath, destPath); err != nil {
			return err
		}
	}
	return nil
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return nil
}
