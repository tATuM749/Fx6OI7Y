// 代码生成时间: 2025-08-09 10:26:29
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/gorilla/mux"
)

// 定义一个Handler函数用于计算哈希值
func hashHandler(w http.ResponseWriter, r *http.Request) {
    // 从URL参数中获取文本
    vars := mux.Vars(r)
    text := vars["text"]

    if text == "" {
        // 如果没有提供文本，返回错误信息
        http.Error(w, "No text provided.", http.StatusBadRequest)
        return
    }

    // 使用SHA256算法计算哈希值
    hash := sha256.Sum256([]byte(text))

    // 将哈希值转换为十六进制字符串
    hashString := hex.EncodeToString(hash[:])

    // 将哈希值返回给客户端
    fmt.Fprintf(w, "The SHA256 hash of '%s' is: %s", text, hashString)
}

func main() {
    // 创建一个新的路由器
    router := mux.NewRouter()

    // 定义路由规则，将路径与处理器函数关联
    router.HandleFunc("/hash/{text}", hashHandler).Methods("GET")

    // 启动HTTP服务器
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
