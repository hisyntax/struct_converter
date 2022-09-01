package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hisyntax/struct-converter/converter"
	"github.com/joho/godotenv"
)

func main() {
	l := log.New(os.Stdout, "json-converter-server ", log.LstdFlags)
	ch := converter.NewConverter(l)

	if err := godotenv.Load(); err != nil {
		l.Println(".env not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := mux.NewRouter()

	router.HandleFunc("/", ch.JsonToStruct).Methods("POST")
	l.Printf("Server running on port %s", port)
	http.ListenAndServe(":"+port, router)
}
