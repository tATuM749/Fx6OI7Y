// 代码生成时间: 2025-09-02 19:59:04
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

// Database configuration
const (
    dbUser     = "your_username"
    dbPassword = "your_password"
    dbName     = "your_dbname"
    host       = "127.0.0.1:3306"
)

// Database connection string
const connectionString = "%s:%s@tcp(%s)/%s?parseTime=True"

// db is a global variable for the database connection
var db *sql.DB

// initDB sets up the database connection
func initDB() error {
    connStr := fmt.Sprintf(connectionString, dbUser, dbPassword, host, dbName)
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        return err
    }
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
    return nil
}

// CloseDB closes the database connection
func closeDB() error {
    return db.Close()
}

// SampleQuery demonstrates how to use parameterized queries to prevent SQL injection
func SampleQuery(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    param := vars["param"]

    query := `SELECT * FROM users WHERE name = ?`
    rows, err := db.Query(query, param)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Handle the query results
    // ...
}

func main() {
    err := initDB()
    if err != nil {
        log.Fatal("Database connection error: ", err)
    }
    defer closeDB()

    router := mux.NewRouter()
    router.HandleFunc("/query/{param}", SampleQuery).Methods("GET")
    
    http.Handle("/", router)
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}