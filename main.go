package main

import (
	"fmt"

	"github.com/zizimoos/gocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("data : %s\n", block.Data)
		fmt.Printf("hash : %s\n", block.Hash)
		fmt.Printf("prevHash : %s\n\n", block.PrevHash)
	}
}
