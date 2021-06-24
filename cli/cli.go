package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/zizimoos/gocoin/explorer"
	"github.com/zizimoos/gocoin/rest"
)

func usage() {
	fmt.Printf("welcome to go-coin \n\n")
	fmt.Printf("please use the following flags : \n\n")
	fmt.Printf("-port=4000 : set the port of the server \n")
	fmt.Printf("-mode=rest   choose between 'html' and 'rest' \n")
	os.Exit(0)
}

func Start() {
	port := flag.Int("port", 4000, "sets the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()
	if len(os.Args) == 1 {
		usage()
	}

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
