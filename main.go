package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from Home")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

	// chain := blockchain.GetBlockchain()
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Third Block")
	// chain.AddBlock("Fourth Block")
	// for _, block := range chain.AllBlocks() {
	// 	fmt.Printf("data : %s\n", block.Data)
	// 	fmt.Printf("hash : %s\n", block.Hash)
	// 	fmt.Printf("prevHash : %s\n\n", block.PrevHash)
	// }
}
