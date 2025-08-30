// 代码生成时间: 2025-08-30 18:30:59
It includes error handling, comments, and documentation to ensure maintainability and scalability.
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
    "runtime/debug"
    "time"

    "github.com/gorilla/websocket"
)

// Define the WebSocket upgrader
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// Websocket connection
type wsConnection struct {
    conn *websocket.Conn
    send chan []byte
}

// MemoryAnalyzer holds the connections and other state
type MemoryAnalyzer struct {
    connections map[*wsConnection]bool
    broadcast   chan []byte

    // Define a lock for managing state safely
    lock sync.Mutex
}

// NewMemoryAnalyzer creates a new MemoryAnalyzer with a broadcast channel
func NewMemoryAnalyzer() *MemoryAnalyzer {
    return &MemoryAnalyzer{
        connections: make(map[*wsConnection]bool),
        broadcast:   make(chan []byte),
    }
}

// Run starts the memory analysis service
func (ma *MemoryAnalyzer) Run() {
    for {
        select {
        case message := <-ma.broadcast:
            ma.lock.Lock()
            for conn := range ma.connections {
                select {
                case conn.send <- message:
                default:
                    delete(ma.connections, conn)
                    conn.conn.Close()
                }
            }
            ma.lock.Unlock()
        }
    }
}

// ServeWs handles WebSocket requests and serves real-time memory usage data
func (ma *MemoryAnalyzer) ServeWs(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    wsConn := &wsConnection{conn: conn, send: make(chan []byte, 256)}
    ma.lock.Lock()
    ma.connections[wsConn] = true
    ma.lock.Unlock()

    // Start the read and write loops
    go ma.handleRead(wsConn)
    go ma.handleWrite(wsConn)
}

// handleRead reads incoming messages from the WebSocket connection
func (ma *MemoryAnalyzer) handleRead(wsConn *wsConnection) {
    defer func() {
        ma.lock.Lock()
        delete(ma.connections, wsConn)
        ma.lock.Unlock()
    }()

    for {
        _, message, err := wsConn.conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        // Handle incoming messages (if any)
    }
}

// handleWrite writes messages to the WebSocket connection
func (ma *MemoryAnalyzer) handleWrite(wsConn *wsConnection) {
    for message := range wsConn.send {
        if err := wsConn.conn.WriteMessage(websocket.TextMessage, message); err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
    }
}

// BroadcastMemoryUsage sends the current memory usage to all connected clients
func (ma *MemoryAnalyzer) BroadcastMemoryUsage() {
    stats := debug.GCStats{
        LastGC:        time.Now(),
       NumGC:         1,
      // Add other relevant memory usage statistics here
    }
    message, err := json.Marshal(stats)
    if err != nil {
        log.Printf("error: %v", err)
        return
    }
    ma.broadcast <- message
}

func main() {
    ma := NewMemoryAnalyzer()
    go ma.Run()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Serve the WebSocket endpoint
        ma.ServeWs(w, r)
    })

    // Start broadcasting memory usage every 5 seconds
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            ma.BroadcastMemoryUsage()
        }
    }
    // Start the HTTP server
    log.Fatal(http.ListenAndServe(":8080", nil))
}
