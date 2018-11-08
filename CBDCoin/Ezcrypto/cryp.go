package Ezcrypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	//"fmt"
)

/*
func main() {
	message := "hello"
	result := Ez256(message)
	fmt.Println([]byte(result))
	//fmt.Println([]byte("21"))
	fmt.Println(bytes.Equal([]byte(result), []byte(message)))
	msg := "2"
	key := "acac"
	key2 := "cac"
	emsg := EzEncrypt(msg, key)
	fmt.Println(emsg)
	fmt.Println(EzDecrypt(emsg, key2))

}
*/
func Ez256(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}

func EzEncrypt(msg string, key string) string {
	result := append([]byte(msg), []byte(key)...)
	return string(result)
}

func EzDecrypt(msg string, key string) string {
	msgbt := []byte(msg)
	keybt := []byte(key)
	if len(msgbt) < len(keybt) {
		return ""
	}
	if bytes.Equal(msgbt[len(msgbt)-len(keybt):], keybt) {
		return string(msgbt[:len(msgbt)-len(keybt)])
	} else {
		return ""
	}

}
