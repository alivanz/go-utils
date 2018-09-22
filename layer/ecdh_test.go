package layer

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"net"
	"sync"
	"testing"

	"github.com/alivanz/go-utils"
)

func TestECDH(t *testing.T) {
	curve := elliptic.P256()
	priv1, x1, y1, _ := elliptic.GenerateKey(curve, rand.Reader)
	priv2, x2, y2, _ := elliptic.GenerateKey(curve, rand.Reader)
	shared1 := ECDHComputeShared(curve, x2, y2, priv1)
	shared2 := ECDHComputeShared(curve, x1, y1, priv2)
	if !bytes.Equal(shared1, shared2) {
		t.Log(hex.EncodeToString(shared1))
		t.Log(hex.EncodeToString(shared2))
		t.Fail()
	}
}

func TestECDHHandshake(t *testing.T) {
	msg1 := []byte("helloww")
	msg2 := []byte("123123jebiofbevgjwuijf4bgrlwnbhejfhweiflewbfyewifuweflbwufweylfwegfygwef3289fowehiufhiuewhfiuh89289hfushfuhsuifhi")

	listener, _ := net.Listen("tcp", "127.0.0.1:0")
	curve := elliptic.P256()
	var shared1, shared2 []byte
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		conn, _ := listener.Accept()
		shared1, _ = ECDHShake(curve, conn)
		wg.Done()
		csec, _ := ECDHLayer(conn)
		buf := bytes.NewBuffer(nil)
		x := utils.NewBinaryWriter(buf)
		x.WriteCompact(msg1)
		x.WriteCompact(msg2)
		csec.Write(buf.Bytes())
	}()
	conn, _ := net.Dial("tcp", listener.Addr().String())
	shared2, _ = ECDHShake(curve, conn)
	wg.Wait()
	if !bytes.Equal(shared1, shared2) {
		t.Log(shared1)
		t.Log(shared2)
		t.Fail()
	}
	csec, _ := ECDHLayer(conn)
	x := utils.NewBinaryReader(csec)
	msg, _ := x.ReadCompact()
	if !bytes.Equal(msg1, msg) {
		t.Fail()
	}
	msg, _ = x.ReadCompact()
	if !bytes.Equal(msg2, msg) {
		t.Fail()
	}
}
