// 代码生成时间: 2025-09-01 19:08:06
package main

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
    "syscall"
    "time"

    "github.com/gorilla/mux"
)

// ProcessManager 定义进程管理器结构
type ProcessManager struct {
    processes map[string]*ProcessInfo
    mutex     sync.Mutex
}

// ProcessInfo 定义进程信息结构
type ProcessInfo struct {
    Pid      int
    Name     string
    Started time.Time
}

// NewProcessManager 创建新的进程管理器实例
func NewProcessManager() *ProcessManager {
    return &ProcessManager{
        processes: make(map[string]*ProcessInfo),
    }
}

// StartProcess 启动一个新进程
func (pm *ProcessManager) StartProcess(name string) (*ProcessInfo, error) {
    cmd := exec.Command(name)
    if err := cmd.Start(); err != nil {
        return nil, err
    }

    info := &ProcessInfo{
        Pid:      cmd.Process.Pid,
        Name:     name,
        Started: time.Now(),
    }
    pm.mutex.Lock()
    defer pm.mutex.Unlock()
    pm.processes[name] = info

    return info, nil
}

// StopProcess 停止指定名称的进程
func (pm *ProcessManager) StopProcess(name string) error {
    pm.mutex.Lock()
    defer pm.mutex.Unlock()

    info, found := pm.processes[name]
    if !found {
        return fmt.Errorf("process not found: %s", name)
    }

    if err := syscall.Kill(info.Pid, syscall.SIGTERM); err != nil {
        return err
    }
    delete(pm.processes, name)
    return nil
}

// ListProcesses 返回当前所有进程的列表
func (pm *ProcessManager) ListProcesses() []*ProcessInfo {
    pm.mutex.Lock()
    defer pm.mutex.Unlock()

    var list []*ProcessInfo
    for _, info := range pm.processes {
        list = append(list, info)
    }
    return list
}

func main() {
    pm := NewProcessManager()
    router := mux.NewRouter()

    // 启动进程的路由
    router.HandleFunc("/start/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]
        info, err := pm.StartProcess(name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "{\"name\":\"%s\", \"pid\":%d}
", name, info.Pid)
    }