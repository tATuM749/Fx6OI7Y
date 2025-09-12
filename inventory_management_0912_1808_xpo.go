// 代码生成时间: 2025-09-12 18:08:06
// inventory_management.go
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
}

// InventoryService handles inventory operations
type InventoryService struct {
    inventory map[string]InventoryItem
}

// NewInventoryService creates a new instance of InventoryService
func NewInventoryService() *InventoryService {
    return &InventoryService{
        inventory: make(map[string]InventoryItem),
    }
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(item InventoryItem) error {
    if _, exists := s.inventory[item.ID]; exists {
        return errors.New("item with this ID already exists")
    }
    s.inventory[item.ID] = item
    return nil
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(item InventoryItem) error {
    if _, exists := s.inventory[item.ID]; !exists {
        return errors.New("item not found")
    }
    s.inventory[item.ID] = item
    return nil
}

// DeleteItem removes an item from the inventory
func (s *InventoryService) DeleteItem(id string) error {
    if _, exists := s.inventory[id]; !exists {
        return errors.New("item not found")
    }
    delete(s.inventory, id)
    return nil
}

// GetItem retrieves an item by its ID
func (s *InventoryService) GetItem(id string) (*InventoryItem, error) {
    item, exists := s.inventory[id]
    if !exists {
        return nil, errors.New("item not found")
    }
    return &item, nil
}

// GetAllItems returns a list of all items in the inventory
func (s *InventoryService) GetAllItems() []InventoryItem {
    var items []InventoryItem
    for _, item := range s.inventory {
        items = append(items, item)
    }
    return items
}

// InventoryHandler handles HTTP requests for inventory operations
type InventoryHandler struct {
    service *InventoryService
}

// NewInventoryHandler creates a new instance of InventoryHandler
func NewInventoryHandler(service *InventoryService) *InventoryHandler {
    return &InventoryHandler{
        service: service,
    }
}

// GetItemHandler handles GET requests to retrieve an item
func (h *InventoryHandler) GetItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    item, err := h.service.GetItem(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

// AddItemHandler handles POST requests to add a new item
func (h *InventoryHandler) AddItemHandler(w http.ResponseWriter, r *http.Request) {
    var newItem InventoryItem
    if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.service.AddItem(newItem); err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newItem)
}

// UpdateItemHandler handles PUT requests to update an item
func (h *InventoryHandler) UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var updatedItem InventoryItem
    if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updatedItem.ID = id
    if err := h.service.UpdateItem(updatedItem); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedItem)
}

// DeleteItemHandler handles DELETE requests to remove an item
func (h *InventoryHandler) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if err := h.service.DeleteItem(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// GetAllItemsHandler handles GET requests to retrieve all items
func (h *InventoryHandler) GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
    items := h.service.GetAllItems()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func main() {
    router := mux.NewRouter()
    service := NewInventoryService()
    handler := NewInventoryHandler(service)

    // Routes for inventory operations
    router.HandleFunc("/inventory/{id}", handler.GetItemHandler).Methods("GET")
    router.HandleFunc("/inventory", handler.AddItemHandler).Methods("POST")
    router.HandleFunc("/inventory/{id}", handler.UpdateItemHandler).Methods("PUT")
    router.HandleFunc("/inventory/{id}", handler.DeleteItemHandler).Methods("DELETE")
    router.HandleFunc("/inventory", handler.GetAllItemsHandler).Methods("GET")

    // Start the HTTP server
    http.ListenAndServe(":8080", router)
}