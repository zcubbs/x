package sops

import (
	"errors"
	"github.com/zcubbs/x/pgp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func generatePGPKeys(name, email, pubKeyPath, privKeyPath string) error {
	// Create a temporary directory for PGP keys
	tempDir, err := ioutil.TempDir("", "pgp_test")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	// Generate PGP keys
	err = pgp.GeneratePGPKeys(name, email, pubKeyPath, privKeyPath)
	if err != nil {
		return err
	}

	return nil
}

func TestEncryptDecrypt(t *testing.T) {
	// Create a temporary directory for PGP keys
	tempDir, err := os.MkdirTemp("", "pgp_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Generate PGP keys
	pubKeyPath := filepath.Join(tempDir, "test_public.key")
	privKeyPath := filepath.Join(tempDir, "test_private.key")
	err = pgp.GeneratePGPKeys("Test User", "test@example.com", pubKeyPath, privKeyPath)
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

	// Get the key ID from the public key
	keyID, err := extractKeyID(pubKeyPath)
	if err != nil {
		t.Fatalf("Failed to extract key ID: %v", err)
	}

	// Encrypt a test message using the public key
	plaintext := "Hello, world!"
	encrypted, err := Encrypt(plaintext, keyID)
	if err != nil {
		t.Errorf("Encrypt returned an error: %v", err)
	}
	if encrypted == plaintext {
		t.Errorf("Encrypt did not encrypt the input: got %v", encrypted)
	}

	t.Log("Encrypted:", encrypted)

	// Decrypt the message using the private key
	decrypted, err := Decrypt(encrypted, privKeyPath)
	if err != nil {
		t.Errorf("Decrypt returned an error: %v", err)
	}
	if decrypted != plaintext {
		t.Errorf("Decrypt did not return the original message: got %v, want %v", decrypted, plaintext)
	}

	t.Log("Decrypted:", decrypted)
}

// extractKeyID extracts the key ID from the public key file.
func extractKeyID(pubKeyPath string) (string, error) {
	// Run gpg --import to add the public key to the keyring
	importCmd := exec.Command("gpg", "--import", pubKeyPath)
	if err := importCmd.Run(); err != nil {
		return "", err
	}

	// Run gpg --list-keys to get the details of the imported key
	listKeysCmd := exec.Command("gpg", "--list-keys", "--with-colons")
	output, err := listKeysCmd.Output()
	if err != nil {
		return "", err
	}

	// Parse the output to find the Key ID
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "pub") {
			// The line format should be something like: "pub:u:4096:1:72D7468F9F459A1C:1613032103:::u:::scESC:"
			cols := strings.Split(line, ":")
			if len(cols) > 4 {
				return cols[4], nil // The fifth column is the Key ID
			}
		}
	}

	return "", errors.New("key ID not found")
}
