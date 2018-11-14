package types

import ()

type Block struct {
	Head      *Blockhead
	Blockhash string
	Txlist    []*Transaction
}

type Blockhead struct {
	Prehash   string
	Height    int
	Stateroot string
	Txroot    string
	Mitter    string
	Nonce     int
}
