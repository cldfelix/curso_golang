package main

import(
	"fmt"
	"net/http"
)


func handleFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Ola wwww</h1>")
}


func main()  {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Iniciando server na porta :3000")
	http.ListenAndServe(":3000", nil)
}