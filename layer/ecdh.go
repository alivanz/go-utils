package layer

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"io"
	"math/big"

	"github.com/alivanz/go-utils"
)

type ecdhlayer struct {
	rw       utils.BinaryReadWriter
	gcm      cipher.AEAD
	rnonce   utils.Nonce96
	wnonce   utils.Nonce96
	buffered *bytes.Buffer
}

func ECDHComputeShared(curve elliptic.Curve, x, y *big.Int, private []byte) []byte {
	x, _ = curve.ScalarMult(x, y, private)
	return x.Bytes()
}
func ECDHShake(curve elliptic.Curve, conn io.ReadWriter) ([]byte, error) {
	private, x1, y1, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	rw := utils.NewBinaryReadWriter(conn)
	go func() {
		rw.WriteCompact(x1.Bytes())
		rw.WriteCompact(y1.Bytes())
	}()
	bx2, _ := rw.ReadCompact()
	by2, _ := rw.ReadCompact()
	x2 := big.NewInt(0).SetBytes(bx2)
	y2 := big.NewInt(0).SetBytes(by2)
	return ECDHComputeShared(curve, x2, y2, private), nil
}

func ECDHLayer(conn io.ReadWriter) (io.ReadWriter, error) {
	shared, err := ECDHShake(elliptic.P256(), conn)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(shared[:32])
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	return &ecdhlayer{utils.NewBinaryReadWriter(conn), aesgcm, utils.Nonce96{}, utils.Nonce96{}, bytes.NewBuffer(nil)}, nil
}

func (layer *ecdhlayer) Read(b []byte) (int, error) {
	if layer.buffered.Len() > 0 {
		return layer.buffered.Read(b)
	} else {
		layer.buffered.Reset()
	}
	ciphertext, err := layer.rw.ReadCompact()
	bnonce := layer.rnonce.Nonce()
	plain, err := layer.gcm.Open(nil, bnonce, ciphertext, nil)
	if err != nil {
		return 0, err
	}
	layer.buffered.Write(plain)
	return layer.buffered.Read(b)
}
func (layer *ecdhlayer) Write(b []byte) (int, error) {
	bnonce := layer.wnonce.Nonce()
	ciphertext := layer.gcm.Seal(nil, bnonce, b, nil)
	err := layer.rw.WriteCompact(ciphertext)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
