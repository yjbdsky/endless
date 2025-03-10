package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yjbdsky/endless"
)

func handler5(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WORLD!"))
}

func main() {
	mux1 := mux.NewRouter()
	mux1.Schemes("https")
	mux1.HandleFunc("/hello", handler5).
		Methods("GET")

	err := endless.ListenAndServeTLS("localhost:4242", "cert.pem", "key.pem", mux1)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 4242 stopped")

	os.Exit(0)
}
