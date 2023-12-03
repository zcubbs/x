package pgp

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratePGPKeys(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "pgp_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	// Cleanup function to remove the temporary directory
	defer os.RemoveAll(tempDir)

	t.Log("tempDir:", tempDir)

	// Define file paths for the public and private keys in the temporary directory
	pubKeyPath := filepath.Join(tempDir, "test_public.key")
	privKeyPath := filepath.Join(tempDir, "test_private.key")

	t.Log("pubKeyPath:", pubKeyPath)
	t.Log("privKeyPath:", privKeyPath)

	// Call the GeneratePGPKeys function
	err = GeneratePGPKeys("Test User", "test@example.com", pubKeyPath, privKeyPath)
	if err != nil {
		t.Fatalf("Failed to generate PGP keys: %v", err)
	}

	// Import the generated PGP keys into the GnuPG keyring
	importCmd := exec.Command("gpg", "--import", pubKeyPath)
	if err := importCmd.Run(); err != nil {
		t.Fatalf("Failed to import public key into GnuPG keyring: %v", err)
	}
	importCmd = exec.Command("gpg", "--allow-secret-key-import", "--import", privKeyPath)
	if err := importCmd.Run(); err != nil {
		t.Fatalf("Failed to import private key into GnuPG keyring: %v", err)
	}

	// Check if public key file exists and is not empty
	if info, err := os.Stat(pubKeyPath); os.IsNotExist(err) || info.Size() == 0 {
		t.Errorf("Public key file was not created or is empty")
	}

	// Check if private key file exists and is not empty
	if info, err := os.Stat(privKeyPath); os.IsNotExist(err) || info.Size() == 0 {
		t.Errorf("Private key file was not created or is empty")
	}
}
