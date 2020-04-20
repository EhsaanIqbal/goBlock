package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//BlockChain is a collection of blocks
type BlockChain struct {
	blocks []*Block
}

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// AddBlock adds blocks to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// DeriveHash derives hash value from the previous hash and data contained in the block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Genesis block
func Genesis() *Block {
	return CreateBlock("Genisus", []byte{})
}

// InitBlockChain is used to initialize the blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("after geniesis")
	chain.AddBlock("after after geniesis")
	chain.AddBlock("after after after geniesis")

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
