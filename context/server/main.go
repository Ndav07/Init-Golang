package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")

	select {
	case <- time.After(time.Second * 5):
		log.Println("Requisição processada com sucesso")
		w.Write([]byte("Requisição processada com sucesso"))
	case <- ctx.Done():
		
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}