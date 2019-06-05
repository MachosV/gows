package main

import (
	"log"
	"mux"
	"net/http"
	_ "views"
)

func main() {
	log.Println("Server up and running 127.0.0.1:8000")
	http.ListenAndServe("127.0.0.1:8000", mux.GetMux())
}
