package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/zizimoos/gocoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(rw, "Hello from Home")
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)

}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
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
