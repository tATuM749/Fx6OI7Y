// 代码生成时间: 2025-09-06 07:18:00
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gorilla/mux"
    "github.com/spf13/viper"
)

// ConfigManager 用于管理配置文件
type ConfigManager struct {
    viper *viper.Viper
}

// NewConfigManager 创建一个新的配置管理器实例
func NewConfigManager(configFile string) (*ConfigManager, error) {
    // 初始化 viper
    v := viper.New()
    v.SetConfigFile(configFile)
    if err := v.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    return &ConfigManager{viper: v}, nil
}

// GetConfig 返回 viper 实例，用于获取配置信息
func (cm *ConfigManager) GetConfig() *viper.Viper {
    return cm.viper
}

func main() {
    // 设置默认的配置文件路径
    configFilePath := "config.yml"
    cm, err := NewConfigManager(configFilePath)
    if err != nil {
        log.Fatalf("Error initializing config manager: %v", err)
    }

    // 获取配置信息
    config := cm.GetConfig()
    // 打印配置信息
    fmt.Printf("Config Data: %+v
", config.AllSettings())

    // 设置路由
    router := mux.NewRouter()
    router.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
        // 获取并返回配置信息
        config := cm.GetConfig()
        fmt.Fprintf(w, "Config Data: %s", config.AllSettings())
    }).Methods("GET")

    // 启动服务器
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
