package pgp

import (
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"os"
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
