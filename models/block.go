package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	PreviousHash string
	Position     int
	Data         *BookCheckout
	TimeStamp    string
	Hash         string
}

func NewGenesisBlock() *Block {
	bc := NewBookCheckout(nil, "")
	bc.IsGenesis = true
	block := &Block{
		Data:      bc,
		Position:  0,
		TimeStamp: time.Now().String(),
	}
	block.GenerateHash()
	return block
}

func (b *Block) GenerateHash() {
	data := fmt.Sprint(b.Position) + b.TimeStamp + b.PreviousHash
	n := sha256.New()
	n.Write([]byte(data))
	b.Hash = fmt.Sprintf("%x", n.Sum(nil))
}

func NewBlock(checkout *BookCheckout, bc *BlockChain) *Block {
	prev := bc.Blocks[len(bc.Blocks)-1]
	b := &Block{
		PreviousHash: prev.Hash,
		Position:     prev.Position + 1,
		TimeStamp:    time.Now().String(),
		Data:         checkout,
	}
	b.GenerateHash()
	return b
}
