package layer

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"testing"

	"github.com/alivanz/go-utils"
)

func TestAESEncrypt(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("hello world")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	var nonce utils.Nonce96

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	bnonce := nonce.Nonce()
	ciphertext := aesgcm.Seal(nil, bnonce, plaintext, nil)
	t.Log(hex.EncodeToString(ciphertext))

	if data, err := aesgcm.Open(nil, bnonce, ciphertext, nil); err != nil {
		t.Log(err)
		t.Fail()
	} else if !bytes.Equal(plaintext, data) {
		t.Log("mismatch")
		t.Fail()
	}
}
