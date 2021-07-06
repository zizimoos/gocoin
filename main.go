package main

import (
	"github.com/zizimoos/gocoin/cli"
	"github.com/zizimoos/gocoin/db"
)

func main() {
	defer db.Close()
	// go explorer.Start(3000)
	// rest.Start(4000)
	// blockchain.Blockchain()
	cli.Start()
	// wallet.Wallet()
}
