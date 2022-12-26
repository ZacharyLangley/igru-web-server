package authentication

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	createKeyFile string
	createKeyType string
	k8sTemplate   *template.Template
)

type k8sTemplateContent struct {
	Name       string
	Namespace  string
	KeyType    string
	PrivateKey string
	PublicKey  string
}

func init() {
	var err error
	k8sTemplate, err = template.New("k8s").Parse(`apiVersion: v1
data:
  {{.Name}}: {{.PrivateKey}}
  {{.Name}}.pub: {{.PublicKey}}
kind: Secret
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
  annotations:
    keyType: {{.KeyType}}
type: Opaque
`)
	if err != nil {
		panic(err)
	}
}

func createKey(cmd *cobra.Command, args []string) error {
	var privBlock pem.Block
	var pubBlock pem.Block
	switch strings.ToLower(createKeyType) {
	case "ed25519":
		privBlock.Type = "PRIVATE KEY"
		pubBlock.Type = "PUBLIC KEY"
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return err
		}
		privBlock.Bytes, err = x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return err
		}
		pubBlock.Bytes, err = x509.MarshalPKIXPublicKey(pub)
		if err != nil {
			return err
		}
	case "rsa":
		privBlock.Type = "PRIVATE KEY"
		pubBlock.Type = "PUBLIC KEY"
		priv, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return err
		}
		privBlock.Bytes, err = x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return err
		}
		pubBlock.Bytes, err = x509.MarshalPKIXPublicKey(priv.PublicKey)
		if err != nil {
			return err
		}
	case "ecdsa":
		privBlock.Type = "PRIVATE KEY"
		pubBlock.Type = "PUBLIC KEY"
		priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return err
		}
		privBlock.Bytes, err = x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return err
		}
		pubBlock.Bytes, err = x509.MarshalPKIXPublicKey(priv.PublicKey)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unsupported key type: %q", createKeyType)
	}

	ext := strings.TrimSpace(strings.ToLower(path.Ext(createKeyFile)))
	if ext == ".yml" || ext == ".yaml" {
		// Create k8s secret manifest
		name := strings.TrimRight(path.Base(createKeyFile), path.Ext(createKeyFile))
		zap.L().Info("Creating k8s secret manifest", zap.String("name", name), zap.String("fileName", createKeyFile))
		manifestFile, err := os.Create(createKeyFile)
		if err != nil {
			return err
		}
		defer manifestFile.Close()
		content := k8sTemplateContent{
			Name:       name,
			Namespace:  "default",
			KeyType:    createKeyType,
			PublicKey:  base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&pubBlock)),
			PrivateKey: base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&privBlock)),
		}
		if err := k8sTemplate.Execute(manifestFile, content); err != nil {
			return err
		}
	} else {
		zap.L().Info("Creating file key pair", zap.String("fileName", createKeyFile))
		privFile, err := os.Create(createKeyFile)
		if err != nil {
			return err
		}
		defer privFile.Close()
		pubFile, err := os.Create(createKeyFile + ".pub")
		if err != nil {
			return err
		}
		defer pubFile.Close()
		if err := pem.Encode(privFile, &privBlock); err != nil {
			return err
		}
		if err := pem.Encode(pubFile, &pubBlock); err != nil {
			return err
		}

		key, err := jwk.ReadFile(createKeyFile)
		if err != nil {
			return fmt.Errorf("failed to read in jwk: %w", err)
		}
		iter := key.Keys(context.Background())
		for iter.Next(context.Background()) {
			pair := iter.Pair().Value
			zap.L().Info("create new pair", zap.Any("pair", pair))
		}
	}
	return nil
}
