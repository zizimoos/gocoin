package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zizimoos/gocoin/blockchain"
)

var templates *template.Template

const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(rw, "Hello from Home")
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)

}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	hander := mux.NewRouter()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	hander.HandleFunc("/", home)
	hander.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), hander))
}
