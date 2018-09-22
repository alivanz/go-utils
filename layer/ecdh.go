package layer

import (
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
	buffered []byte
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
	rnonce := big.NewInt(0)
	wnonce := big.NewInt(0)
	rnonce.SetBit(rnonce, 88, 1)
	wnonce.SetBit(wnonce, 88, 1)
	return &ecdhlayer{utils.NewBinaryReadWriter(conn), aesgcm, utils.Nonce96{}, utils.Nonce96{}, []byte{}}, nil
}

func (layer *ecdhlayer) Read(b []byte) (int, error) {
	if len(layer.buffered) > 0 {
		n := copy(b, layer.buffered)
		layer.buffered = layer.buffered[n:]
		return n, nil
	}
	ciphertext, err := layer.rw.ReadCompact()
	bnonce := layer.rnonce.Nonce()
	plain, err := layer.gcm.Open(nil, bnonce, ciphertext, nil)
	if err != nil {
		return 0, err
	}
	n := copy(b, plain)
	layer.buffered = plain[n:]
	return len(b), nil
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
