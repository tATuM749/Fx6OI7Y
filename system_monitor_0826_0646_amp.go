// 代码生成时间: 2025-08-26 06:46:08
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
    "runtime"
    "github.com/gorilla/mux"
)

// SystemInfo holds the system performance data
type SystemInfo struct {
    UpTime             string    `json:"upTime"`
    TotalProcesses     int       `json:"totalProcesses"`
    RunningProcesses  int       `json:"runningProcesses"`
    CpuPercent         float64   `json:"cpuPercent"`
    MemoryPercent      float64   `json:"memoryPercent"`
    DiskPercent       float64   `json:"diskPercent"`
}

// GetSystemInfo fetches system performance data
func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
    // Retrieve system uptime
    upTime := time.Since(time.Now())
    upTimeStr := fmt.Sprintf("%v", upTime)

    // Retrieve total and running processes
    totalProcesses, err := os.Getpid()
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    runningProcesses, err := os.Getpid()
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Retrieve CPU usage percentage
    cpuPercent := 0.0 // Placeholder for actual CPU usage calculation

    // Retrieve memory usage percentage
    memStats := runtime.MemStats{}
    runtime.ReadMemStats(&memStats)
    memoryPercent := float64(memStats.Alloc) / float64(runtime.RuntimeMemStats{}.Alloc) * 100

    // Retrieve disk usage percentage
    diskPercent := 0.0 // Placeholder for actual disk usage calculation

    // Create and send SystemInfo struct
    systemInfo := SystemInfo{
        UpTime:             upTimeStr,
        TotalProcesses:     totalProcesses,
        RunningProcesses:  runningProcesses,
        CpuPercent:         cpuPercent,
        MemoryPercent:      memoryPercent,
        DiskPercent:       diskPercent,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(systemInfo)
}

func main() {
    router := mux.NewRouter()
    // Define the route for fetching system info
    router.HandleFunc("/system", GetSystemInfo).Methods("GET")

    // Start the server
    fmt.Println("Starting system monitor tool...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
