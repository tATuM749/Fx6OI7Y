// 代码生成时间: 2025-09-14 14:26:26
package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

// HandlerFunc is a type that defines the signature of a request handler function.
// It takes an http.ResponseWriter and an *http.Request as arguments.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// NotFoundHandler is a request handler for 404 Not Found responses.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}

// HomeHandler is the handler for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Respond with a simple message in the HTTP response.
    fmt.Fprintf(w, "Welcome to the home page!")
}

// AboutHandler is the handler for the about page.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
    // Respond with a simple message in the HTTP response.
    fmt.Fprintf(w, "Welcome to the about page!")
}

// main is the entry point of the program. It sets up the HTTP server and routes.
func main() {
    // Create a new router.
    router := mux.NewRouter()

    // Define routes with their respective handlers.
    router.HandleFunc("/", HomeHandler).Methods("GET")
    router.HandleFunc("/about", AboutHandler).Methods("GET\)

    // Add a custom not found handler for all undefined routes.
    router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

    // Start the server on port 8080.
    log.Println("Starting server on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}
