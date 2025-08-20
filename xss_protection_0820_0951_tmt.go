// 代码生成时间: 2025-08-20 09:51:34
// Package main provides a basic example of how to protect against XSS attacks using the Gorilla framework.
package main

import (
    "net/http"
    "html"
    "log"
    "github.com/gorilla/mux"
)

// HTMLEscape is a middleware that escapes HTML to prevent XSS attacks.
func HTMLEscape(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        for key, value := range r.URL.Query() {
            value = html.EscapeString(value[0])
            r.URL.RawQuery = strings.ReplaceAll(r.URL.RawQuery, value[0], value)
        }
        h.ServeHTTP(w, r)
    })
}

// IndexHandler is the handler for the index route. It shows a simple page with user input.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    userInput := r.FormValue("userInput")
    if userInput == "" {
        userInput = "No user input provided."
    }

    // Sanitize the user input to prevent XSS.
    sanitizedInput := html.EscapeString(userInput)

    // Render the template with the sanitized input.
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, map[string]string{"UserInput": sanitizedInput})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler).Methods("GET", "POST")

    // Apply the HTMLEscape middleware to all routes.
    http.Handle("/", HTMLEscape{r})

    log.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080