package main

import (
	. "./types"
	//"encoding/json"
	"fmt"
	//"github.com/oj01ol/CBDCoin/CBDCoin/Ezcrypto"
)

func main() {
	k := NewKey("seed")
	m := NewKey("ha")
	txa1 := &Txatom{
		Address: k.Pubkey,
		Mitter:  m.Pubkey,
		Amount:  10,
	}
	txa2 := &Txatom{
		Address: "bf",
		Mitter:  m.Pubkey,
		Amount:  9,
	}
	tdata := &Txdata{
		From: []*Txatom{txa1},
		To:   []*Txatom{txa2},
	}
	sig1 := &Sigatom{
		Address: k.Pubkey,
		Sig:     k.Sig(tdata.Hash()),
	}
	sige := &Sigatom{
		Address: m.Pubkey,
		Sig:     m.Sig(tdata.Hash()),
	}
	tx := &Transaction{
		Data:        tdata,
		Txhash:      tdata.Hash(),
		Txsig:       []*Sigatom{sig1},
		Txevidences: []*Sigatom{sige},
		Txcommits:   []*Sigatom{sige},
	}
	fmt.Println(tx.Check())
	fmt.Println(tx.Txhash)
	fmt.Println(tx.Checksig())
	/*
		k := NewKey("seed")
		fmt.Println(k)
		sig := k.Sig("hello")
		fmt.Println(sig)
		fmt.Println(VerifySig("hello", sig, k.Pubkey))
	*/
}
