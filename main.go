package main

import (
	"fmt"
	"strconv"

	"github.com/ehsaaniqbal/goBlock/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("after geniesis")
	chain.AddBlock("after after geniesis")
	chain.AddBlock("after after after geniesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
