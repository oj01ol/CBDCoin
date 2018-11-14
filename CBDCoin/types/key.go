package types

import (
	"bytes"
	"crypto/rand"
	//"fmt"
	"github.com/oj01ol/CBDCoin/CBDCoin/Ezcrypto"
)

type Key struct {
	Pubkey string
	Prikey string
	Seed   string
}

func NewKey(seed string) *Key {
	if seed == "" {
		b := make([]byte, 20)
		rand.Read(b)
		seed = string(b)
	}
	k := Key{
		Pubkey: Ezcrypto.Ez256(seed),
		Prikey: Ezcrypto.Ez256(seed),
		Seed:   seed,
	}
	return &k
}

func (k *Key) Sig(msg string) string {
	return Ezcrypto.EzEncrypt(msg, k.Prikey)
}

func VerifySig(msg string, sig string, pubkey string) bool {
	smsg := Ezcrypto.EzDecrypt(sig, pubkey)
	return bytes.Equal([]byte(msg), []byte(smsg))
}

/*
func main() {
	k := NewKey("seed")
	fmt.Println(k)
	sig := k.Sig("hello")
	fmt.Println(sig)
	fmt.Println(VerifySig("hello", sig, k.Pubkey))

}
*/
