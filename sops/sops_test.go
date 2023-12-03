package sops

import (
	"github.com/zcubbs/x/pgp"
	"os"
	"path/filepath"
	"testing"
)

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
	err = pgp.ImportPublicKey(pubKeyPath, privKeyPath)
	if err != nil {
		t.Fatalf("Failed to import public key: %v", err)
	}

	// Get the key ID from the public key
	keyID, err := pgp.ExtractKeyID(pubKeyPath)
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
