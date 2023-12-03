package pgp

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"os"
	"os/exec"
	"strings"
	"time"
)

// GeneratePGPKeys generates a PGP public and private key pair and saves them to the given file paths.
func GeneratePGPKeys(name, email, pubKeyPath, privKeyPath string) error {
	// Create the key pair
	entity, err := openpgp.NewEntity(name, "PGP key pair", email, nil)
	if err != nil {
		return err
	}

	// Set validity of the key
	entity.PrimaryKey.CreationTime = time.Now()
	for _, identity := range entity.Identities {
		identity.SelfSignature.CreationTime = entity.PrimaryKey.CreationTime
		// identity.SelfSignature.KeyLifetimeSecs = ... (optional, set if needed)
	}
	for _, subkey := range entity.Subkeys {
		subkey.Sig.CreationTime = entity.PrimaryKey.CreationTime
		// subkey.Sig.KeyLifetimeSecs = ... (optional, set if needed)
	}

	// Save the private key
	privKeyFile, err := os.Create(privKeyPath)
	if err != nil {
		return err
	}
	defer privKeyFile.Close()
	privWriter, err := armor.Encode(privKeyFile, openpgp.PrivateKeyType, nil)
	if err != nil {
		return err
	}
	defer privWriter.Close()
	err = entity.SerializePrivate(privWriter, nil)
	if err != nil {
		return err
	}

	// Save the public key
	pubKeyFile, err := os.Create(pubKeyPath)
	if err != nil {
		return err
	}
	defer pubKeyFile.Close()
	pubWriter, err := armor.Encode(pubKeyFile, openpgp.PublicKeyType, nil)
	if err != nil {
		return err
	}
	defer pubWriter.Close()
	err = entity.Serialize(pubWriter)
	if err != nil {
		return err
	}

	return nil
}

// ExtractKeyID extracts the key ID from a PGP public key file.
func ExtractKeyID(pubKeyPath string) (string, error) {
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

// ImportPublicKey imports a PGP public key into the GnuPG keyring.
func ImportPublicKey(pubKeyPath, privKeyPath string) error {
	// Import the generated PGP keys into the GnuPG keyring
	importCmd := exec.Command("gpg", "--import", pubKeyPath)
	if err := importCmd.Run(); err != nil {
		return fmt.Errorf("failed to import public key into GnuPG keyring: %v", err)
	}
	importCmd = exec.Command("gpg", "--allow-secret-key-import", "--import", privKeyPath)
	if err := importCmd.Run(); err != nil {
		return fmt.Errorf("failed to import private key into GnuPG keyring: %v", err)
	}

	return nil
}
