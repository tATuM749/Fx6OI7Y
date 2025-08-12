// 代码生成时间: 2025-08-13 05:06:57
package main

import (
    "fmt"
    "net/http"
    "html/template"
    "log"

    "github.com/gorilla/mux"
)

// HTML templates
const (
    indexTemplate = `<html>
<head>
<title>Responsive Design Demo</title>
<meta name="viewport" content="width=device-width, initial-scale=1"/>
<style>
body { margin: 0; font-family: Arial, sans-serif; }
header { background: #333; color: white; padding: 10px 0; text-align: center; }
nav { display: flex; justify-content: center; padding: 10px 0; }
nav a { color: #333; text-decoration: none; padding: 10px 15px; }
nav a:hover { background: #f4f4f4; }
@media (max-width: 600px) {
    nav { flex-direction: column; }
    nav a { padding: 10px; }
}
</style>
</head>
<body>
<header>Responsive Design Demo</header>
<nav>
{{range .Links}}
<a href="{{.URL}}">{{.Name}}</a>
{{end}}
</nav>
</body>
</html>`
)

// Link represents a navigation link
type Link struct {
    Name string
    URL  string
}

// IndexData holds data for the index template
type IndexData struct {
    Links []Link
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")

    // Define navigation links
    links := []Link{
        {
            Name: "Home",
            URL:  "/",
        },
        {
            Name: "About",
            URL:  "/about",
        },
        {
            Name: "Contact",
            URL:  "/contact",
        },
    }

    // Create a template and parse the index template
    tmpl := template.New("index")
    tmpl, err := tmpl.ParseGlob("indexTemplate")
    if err != nil {
        log.Fatal(err)
    }

    // Start the server
    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}

// indexHandler handles requests to the root path and serves the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Prepare data for the template
    data := IndexData{
        Links: []Link{
            {
                Name: "Home",
                URL:  "/",
            },
            {
                Name: "About",
                URL:  "/about",
            },
            {
                Name: "Contact",
                URL:  "/contact",
            },
        },
    }

    // Execute the template with the data
    if err := tmpl.ExecuteTemplate(w, "index", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
