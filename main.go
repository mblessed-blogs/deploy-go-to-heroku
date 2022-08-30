package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// home handler is unexported due to no other package requiring the handler
func home(w http.ResponseWriter, r *http.Request) {
	response := "Home page"
	json.NewEncoder(w).Encode(response)
}

// portNumber
const portNumber = ":8000"

func main() {
	// create a new router
	router := chi.NewRouter()
	router.HandleFunc("/", home)

	// create our server
	srv := &http.Server{
		Addr: portNumber,
		Handler: router,
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}