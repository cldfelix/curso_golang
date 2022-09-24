package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Ola wwww</h1>")
}

func main() {
	//mux := http.NewServeMux()
	time.Sleep(1 * time.Second)
	http.HandleFunc("/", handleFunc)
	fmt.Println("Iniciando server na porta :3000")
	//fmt.Fprintln(os.Stdout, "Iniciando server na porta :3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
