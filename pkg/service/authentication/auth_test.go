package authentication

import (
	"encoding/base64"
	"testing"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
)

func FuzzPassword(f *testing.F) {
	f.Add("eric")
	f.Add("zach")
	f.Add("")
	f.Add("fddfdadsfkjasdflkjadsfnlzvilfeio")
	f.Add(".")
	f.Fuzz(func(t *testing.T, a string) {
		hash, salt, err := generateCred(a)
		if err != nil {
			t.Error(err)
		}
		if _, err := base64.StdEncoding.DecodeString(hash); err != nil {
			t.Error(err)
		}
		if _, err := base64.StdEncoding.DecodeString(salt); err != nil {
			t.Error(err)
		}
		user := models.User{
			Hash: hash,
			Salt: salt,
		}
		if !checkHash(user, a) {
			t.Error("failed to check hash")
		}
	})
}
