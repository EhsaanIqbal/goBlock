package main

import (
	"fmt"
)

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
