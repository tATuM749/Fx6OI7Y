// 代码生成时间: 2025-08-01 20:58:26
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Data represents the structure of the JSON data to be converted.
type Data struct {
    Value string `json:"value"`
}

// ConvertJSON is the handler function that will be called for the conversion endpoint.
// It takes a JSON payload, converts it, and returns the same payload as the response.
func ConvertJSON(w http.ResponseWriter, r *http.Request) {
    var data Data
    // Decode the incoming JSON request into our Data struct.
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        // If there is an error decoding the JSON, send a 400 Bad Request response.
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Encode the data back into JSON to send as the response.
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(data); err != nil {
        // If there is an error encoding the JSON, log the error.
        log.Fatalf("Error encoding JSON: %s", err)
    }
}

func main() {
    // Create a new router using the gorilla/mux package.
    router := mux.NewRouter()
    // Define a route for the JSON conversion endpoint.
    router.HandleFunc("/convert", ConvertJSON).Methods("POST")

    // Start the server using the router.
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
