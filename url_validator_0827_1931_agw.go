// 代码生成时间: 2025-08-27 19:31:40
package main

import (
    "fmt"
    "net/http"
    "net/url"
    "log"
    "github.com/gorilla/mux"
)

// URLValidatorService 结构体，用于封装URL验证的逻辑
type URLValidatorService struct {
    // 可以添加配置参数，例如超时时间等
}

// NewURLValidatorService 创建并返回一个新的URLValidatorService实例
func NewURLValidatorService() *URLValidatorService {
    return &URLValidatorService{}
}

// ValidateURL 验证给定的URL是否有效
func (s *URLValidatorService) ValidateURL(rawURL string) (bool, error) {
    // 解析URL
    parsedURL, err := url.ParseRequestURI(rawURL)
    if err != nil {
        return false, fmt.Errorf("failed to parse URL: %w", err)
    }

    // 创建一个新的HTTP客户端
    client := &http.Client{}
    // 发送HEAD请求检查URL是否可达
    resp, err := client.Head(parsedURL.String())
    if err != nil {
        return false, fmt.Errorf("failed to reach URL: %w", err)
    }
    defer resp.Body.Close()

    // 检查响应状态码
    if resp.StatusCode == http.StatusOK {
        return true, nil
    } else {
        return false, fmt.Errorf("URL returned status code %d", resp.StatusCode)
    }
}

// setupRoutes 设置路由并启动服务器
func setupRoutes(r *mux.Router, service *URLValidatorService) {
    // 为URL验证设置路由
    r.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
        // 从请求中获取URL参数
        url := r.FormValue("url")
        if url == "" {
            http.Error(w, "URL parameter is required", http.StatusBadRequest)
            return
        }

        // 使用服务验证URL
        valid, err := service.ValidateURL(url)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // 返回验证结果
        if valid {
            fmt.Fprintf(w, "The URL is valid.")
        } else {
            fmt.Fprintf(w, "The URL is not valid.")
        }
    }
)
}

func main() {
    // 创建一个新的路由
    router := mux.NewRouter()
    // 创建URL验证服务
    service := NewURLValidatorService()

    // 设置路由
    setupRoutes(router, service)

    // 启动服务器
    log.Println("Starting server on port 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}