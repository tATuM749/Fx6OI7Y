// 代码生成时间: 2025-08-31 23:42:01
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

// User represents a user with a username and password.
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// ErrorResponse defines the structure for the error response.
type ErrorResponse struct {
    Error string `json:"error"`
}

// loginHandler handles the login requests.
func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Only allow POST requests.
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode the user data from the request body.
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Perform login validation (for simplicity, using hardcoded credentials).
    if user.Username != "admin" || user.Password != "password" {
        // Respond with an error message.
        resp := ErrorResponse{Error: "Invalid credentials"}
        json.NewEncoder(w).Encode(resp)
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // Respond with a success message.
    response := map[string]string{"message": "Login successful"}
    json.NewEncoder(w).Encode(response)
    w.WriteHeader(http.StatusOK)
}

func main() {
    // Create a new router instance.
    r := mux.NewRouter()

    // Register the login handler.
    r.HandleFunc("/login", loginHandler).Methods("POST")

    // Start the server on port 8080.
    http.ListenAndServe(":8080", r)
}
