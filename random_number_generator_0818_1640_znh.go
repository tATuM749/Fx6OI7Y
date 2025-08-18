// 代码生成时间: 2025-08-18 16:40:28
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/gorilla/mux"
)

// RandomNumberGenerator struct represents the random number generator service
type RandomNumberGenerator struct {
    // No fields required for this simple service
}

// NewRandomNumberGenerator initializes and returns a new instance of RandomNumberGenerator
func NewRandomNumberGenerator() *RandomNumberGenerator {
    return &RandomNumberGenerator{}
}

// GenerateRandomNumber generates a random number between min and max values
func (rng *RandomNumberGenerator) GenerateRandomNumber(min, max int) (int, error) {
    if min > max {
        return 0, fmt.Errorf("min value cannot be greater than max value")
    }
    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min, nil
}

// RandomNumberHandler handles the HTTP request to generate a random number
func RandomNumberHandler(rng *RandomNumberGenerator) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        min := 1 // Default min value
        max := 100 // Default max value

        // Extract query parameters for min and max, if provided
        var err error
        if minStr := r.URL.Query().Get("min"); minStr != "" {
            min, err = strconv.Atoi(minStr)
            if err != nil {
                http.Error(w, "Invalid min value", http.StatusBadRequest)
                return
            }
        }

        if maxStr := r.URL.Query().Get("max"); maxStr != "" {
            max, err = strconv.Atoi(maxStr)
            if err != nil {
                http.Error(w, "Invalid max value", http.StatusBadRequest)
                return
            }
        }

        randomNumber, err := rng.GenerateRandomNumber(min, max)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Write the random number as a JSON response
        response := struct{
            Number int `json:"number"`
        }{
            Number: randomNumber,
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func main() {
    // Initialize the random number generator service
    rng := NewRandomNumberGenerator()

    // Create a new router
    router := mux.NewRouter()

    // Define the route for generating a random number
    router.HandleFunc("/random", RandomNumberHandler(rng)).Methods("GET")

    // Start the HTTP server
    fmt.Println("Server is running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
