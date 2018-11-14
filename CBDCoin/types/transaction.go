package types

import (
	"bytes"
	"encoding/json"
	"github.com/oj01ol/CBDCoin/CBDCoin/Ezcrypto"
)

const Maxint = 10000000000000000

type Transaction struct {
	Data        *Txdata
	Txhash      string
	Txsig       []*Sigatom
	Txevidences []*Sigatom
	Txcommits   []*Sigatom
	//status      int
}

type Txdata struct {
	From []*Txatom
	To   []*Txatom
}

type Txatom struct {
	Address string
	Mitter  string
	Amount  int
}

type Sigatom struct {
	Address string
	Sig     string
}

func NewTx(from []*Txatom, to []*Txatom) *Transaction {
	txdata := &Txdata{
		From: from,
		To:   to,
	}
	tx := &Transaction{
		Data:   txdata,
		Txhash: txdata.Hash(),
	}
	return tx
}

func (tx *Transaction) Addsig(newsig *Sigatom) *Transaction {
	for _, siglist := range tx.Txsig {
		if bytes.Equal([]byte(newsig.Address), []byte(siglist.Address)) {
			return tx
		}
	}
	tx.Txsig = append(tx.Txsig, newsig)
	return tx
}

func (tx *Transaction) Addevi(newsig *Sigatom) *Transaction {
	for _, siglist := range tx.Txevidences {
		if bytes.Equal([]byte(newsig.Address), []byte(siglist.Address)) {
			return tx
		}
	}
	tx.Txevidences = append(tx.Txevidences, newsig)
	return tx
}

func (tx *Transaction) Addcom(newsig *Sigatom) *Transaction {
	for _, siglist := range tx.Txcommits {
		if bytes.Equal([]byte(newsig.Address), []byte(siglist.Address)) {
			return tx
		}
	}
	tx.Txcommits = append(tx.Txcommits, newsig)
	return tx
}

func (tdata *Txdata) CheckAmount() bool {
	in := 0
	out := 0
	for _, inatom := range tdata.From {
		if inatom.Amount <= 0 || inatom.Amount >= Maxint {
			return false
		}
		insum := in + inatom.Amount
		if insum <= in {
			return false
		}
		in = insum
	}
	for _, outatom := range tdata.To {
		if outatom.Amount <= 0 || outatom.Amount >= Maxint {
			return false
		}
		outsum := out + outatom.Amount
		if outsum <= out {
			return false
		}
		out = outsum
	}
	return in >= out

}

func (tdata *Txdata) Hash() string {
	txbyt, _ := json.Marshal(tdata)
	return Ezcrypto.Ez256(string(txbyt))

}

func (tx *Transaction) Checksig() bool {
	needlist := make(map[string]bool)
	for _, inatom := range tx.Data.From {
		needlist[inatom.Address] = false
	}
	for _, sigatom := range tx.Txsig {
		needlist[sigatom.Address] = VerifySig(tx.Txhash, sigatom.Sig, sigatom.Address)
	}
	for _, ok := range needlist {
		if !ok {
			return false
		}
	}
	return true
}

func (tx *Transaction) Checkevi() bool {
	needlist := make(map[string]bool)
	for _, inatom := range tx.Data.From {
		needlist[inatom.Mitter] = false
	}
	for _, sigatom := range tx.Txevidences {
		needlist[sigatom.Address] = VerifySig(tx.Txhash, sigatom.Sig, sigatom.Address)
	}
	for _, ok := range needlist {
		if !ok {
			return false
		}
	}
	return true
}

func (tx *Transaction) Checkcom() bool {
	needlist := make(map[string]bool)
	for _, outatom := range tx.Data.To {
		needlist[outatom.Mitter] = false
	}
	for _, sigatom := range tx.Txcommits {
		needlist[sigatom.Address] = VerifySig(tx.Txhash, sigatom.Sig, sigatom.Address)
	}
	for _, ok := range needlist {
		if !ok {
			return false
		}
	}
	return true
}

func (tx *Transaction) Check() int {
	if !tx.Data.CheckAmount() {
		return 1 //Amount error
	}
	if tx.Txhash != tx.Data.Hash() {
		return 2 //data hash error
	}
	if !tx.Checksig() {
		return 3 //sig error
	}
	if !tx.Checkevi() {
		return 4
	}
	if !tx.Checkcom() {
		return 5
	}

	return 0 //pass

}
