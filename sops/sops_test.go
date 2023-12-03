package sops

import (
	"testing"
)

// TestEncrypt tests the Encrypt function
func TestEncrypt(t *testing.T) {
	plaintext := "Hello, world!"
	encrypted, err := Encrypt(plaintext)
	if err != nil {
		t.Errorf("Encrypt returned an error: %v", err)
	}
	if encrypted == plaintext {
		t.Errorf("Encrypt did not encrypt the input: got %v", encrypted)
	}
}

// TestDecrypt tests the Decrypt function
func TestDecrypt(t *testing.T) {
	plaintext := "Hello, world!"
	encrypted, err := Encrypt(plaintext)
	if err != nil {
		t.Errorf("Encrypt returned an error: %v", err)
	}
	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Errorf("Decrypt returned an error: %v", err)
	}
	if decrypted != plaintext {
		t.Errorf("Decrypt did not return the original message: got %v, want %v", decrypted, plaintext)
	}
}
