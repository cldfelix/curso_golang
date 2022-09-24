package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Ola wwww</h1>")
}

func contacthanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>tela contato</h1>")
}

/*
	func pathHandler(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			handleFunc(w, r)
		case "/contato":
			contacthanler(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
*/
type Route struct{}

func (router Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleFunc(w, r)
	case "/contato":
		contacthanler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}

func main() {
	//mux := http.NewServeMux()
	time.Sleep(1 * time.Second)
	//http.HandleFunc("/", pathHandler)
	var router Route
	fmt.Println("Iniciando server na porta :3000")
	//fmt.Fprintln(os.Stdout, "Iniciando server na porta :3000")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
