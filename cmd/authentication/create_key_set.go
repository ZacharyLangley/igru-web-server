package authentication

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/spf13/cobra"
)

var (
	createKeySetFile string
	createKeySetType string
)

func getKeySet(filename string) (jwk.Set, error) {
	// Open file
	f, err := os.Open(createKeySetFile)
	if err != nil {
		if os.IsNotExist(err) {
			return jwk.NewSet(), nil
		}
		return nil, fmt.Errorf("failed to open key set file: %w", err)
	}
	defer f.Close()
	set, err := jwk.ReadFile(createKeySetFile)
	if err != nil {
		return nil, fmt.Errorf("invalid key set file format: %w", err)
	}
	return set, nil
}

func createKeySet(cmd *cobra.Command, args []string) error {
	set, err := getKeySet(createKeySetFile)
	if err != nil {
		return err
	}
	var publicKey crypto.PublicKey
	var privateKey crypto.PrivateKey
	switch strings.ToLower(createKeyType) {
	case "ed25519":
		publicKey, privateKey, err = ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return err
		}
	case "rsa":
		priv, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return err
		}
		privateKey = priv
		publicKey = priv.PublicKey
	case "ecdsa":
		priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return err
		}
		privateKey = priv
		publicKey = priv.PublicKey
	default:
		return fmt.Errorf("Unsupported key type: %q", createKeyType)
	}

	// Add new key to set
	pubKey, err := jwk.FromRaw(publicKey)
	if err != nil {
		return err
	}
	jwk.AssignKeyID(pubKey)
	set.AddKey(pubKey)
	privKey, err := jwk.FromRaw(privateKey)
	if err != nil {
		return err
	}
	jwk.AssignKeyID(privKey)
	set.AddKey(privKey)

	f, err := os.Create(createKeySetFile)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(set)
}
