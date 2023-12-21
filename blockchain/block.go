package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     		[]byte
	Transactions	[]*Transaction
	PrevHash 		[]byte
	Nonce    		int
}

func NewBlock(txs []*Transaction, prevHash []byte) *Block {
	b := &Block{
		Hash:     		[]byte{},
		Transactions: 	txs,
		PrevHash: 		prevHash,
		Nonce:    		0,
	}
	pow := NewPOW(b)
	nonce, hash := pow.Run()
	b.Hash = hash[:]
	b.Nonce = nonce
	return b
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Serialize())
	}
	tree := NewMerkleTree(txHashes)

	return tree.RootNode.Data
}

func Genesis(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}

	return &block
}