package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"study-app/db_stuff"
	"study-app/handlers"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")
	io.WriteString(w, "This is my website bitch!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /hello request")
	io.WriteString(w, "Hello, mufugga!")
}

func main() {
	db_stuff.Init()
	handlers.Mux = http.NewServeMux()
	handlers.Mux.HandleFunc("/", getRoot)
	handlers.Mux.HandleFunc("/hello", getHello)
	handlers.Mux.HandleFunc("/fart", handlers.GetFart)
	handlers.Mux.HandleFunc("/app/collections", handlers.GetCollections)
	handlers.Mux.HandleFunc("/app/collections/", handlers.GetCollection)

	fmt.Println("Starting a stinky server right about.... NOW!")
	err := http.ListenAndServe(":3333", handlers.Mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server done been closed yo...")
	} else if err != nil {
		log.Fatal("Squeaky fart...")
		os.Exit(1)
	}
}
