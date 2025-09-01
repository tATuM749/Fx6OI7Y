// 代码生成时间: 2025-09-01 08:53:30
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
)

// setupRouter initializes a router as defined by the routes
func setupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Test endpoint works.")
    })
    return r
}

// TestIntegration is an integration test for the HTTP server
func TestIntegration(t *testing.T) {
    // Setup the router
    router := setupRouter()
    server := httptest.NewServer(router)
    defer server.Close()

    // Make a GET request to the test endpoint
    resp, err := http.Get(server.URL + "/test")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    // Check the status code and the body of the response
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
    }

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Errorf("Expected no error reading body, got %v", err)
    }
    defer resp.Body.Close()

    // Check the body content
    expectedBody := "Test endpoint works."
    if string(body) != expectedBody {
        t.Errorf("Expected body %q, got %q", expectedBody, string(body))
    }
}
