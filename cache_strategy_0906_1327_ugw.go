// 代码生成时间: 2025-09-06 13:27:36
// cache_strategy.go

package main

import (
    "fmt"
    "net/http"
    "time"
    "golang.org/x/time/rate"
    "github.com/gorilla/mux"
)

// CacheKey defines the key to store the cached value.
type CacheKey string

// CacheItem holds the cached item with the expiration date.
type CacheItem struct {
    Value      interface{}
    Expiration time.Time
}

// Cache is a simple in-memory cache with expiration.
type Cache struct {
    cache       map[CacheKey]CacheItem
    expiration  time.Duration
    requestRate *rate.Limiter
}

// NewCache creates a new cache with the given expiration duration and request rate limit.
func NewCache(expiration time.Duration, requestRate float64) *Cache {
    return &Cache{
        cache:       make(map[CacheKey]CacheItem),
        expiration:  expiration,
        requestRate: rate.NewLimiter(requestRate, 1),
    }
}

// Set sets the value for a specific key in the cache with expiration.
func (c *Cache) Set(key CacheKey, value interface{}, duration time.Duration) {
    c.cache[key] = CacheItem{
        Value:      value,
        Expiration: time.Now().Add(duration),
    }
}

// Get retrieves a value from the cache based on the given key.
func (c *Cache) Get(key CacheKey) (interface{}, error) {
    if item, exists := c.cache[key]; exists && time.Now().Before(item.Expiration) {
        // Check if the request is allowed based on the rate limit.
        if c.requestRate.Allow() {
            return item.Value, nil
        }
        return nil, fmt.Errorf("request rate limit exceeded")
    }
    return nil, fmt.Errorf("cache item not found or expired")
}

// Clear removes the cached item based on the given key.
func (c *Cache) Clear(key CacheKey) {
    delete(c.cache, key)
}

// StartCacheService starts the HTTP server with caching middleware.
func StartCacheService() {
    r := mux.NewRouter()
    c := NewCache(5*time.Minute, 1.0) // 5 minute expiration, rate limit of 1 request per second
    
    r.HandleFunc("/cache/{key}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        key := CacheKey(vars["key"])
        value, err := c.Get(key)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        
        // Respond with the cached value.
        fmt.Fprintln(w, value)
    }).Methods("GET")
    
    r.HandleFunc("/cache/{key}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        key := CacheKey(vars["key"])
        v := r.FormValue("value")
        duration, _ := time.ParseDuration(r.FormValue("duration") + "m")
        c.Set(key, v, duration)
        fmt.Fprintln(w, "Cached value set")
    }).Methods("POST")
    
    r.HandleFunc("/cache/{key}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        key := CacheKey(vars["key"])
        c.Clear(key)
        fmt.Fprintln(w, "Cached value cleared")
    }).Methods("DELETE\)
    
    // Start the server on the specified port.
    http.ListenAndServe(":8080", r)
}

func main() {
    StartCacheService()
}