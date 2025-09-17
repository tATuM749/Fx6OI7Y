// 代码生成时间: 2025-09-17 18:31:44
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// 文件重命名结构
type RenameRule struct {
    SearchPattern string // 要搜索的模式
    ReplaceWith  string // 替换模式
}

// 文件重命名函数
func RenameFiles(directory string, rules []RenameRule) error {
    files, err := os.ReadDir(directory)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() { // 只处理文件
            for _, rule := range rules {
                oldName := file.Name()
                newName := strings.ReplaceAll(oldName, rule.SearchPattern, rule.ReplaceWith)
                if oldName != newName {
                    // 构建完整的文件路径
                    oldPath := filepath.Join(directory, oldName)
                    newPath := filepath.Join(directory, newName)

                    // 重命名文件
                    if err := os.Rename(oldPath, newPath); err != nil {
                        return fmt.Errorf("failed to rename file from %s to %s: %w", oldPath, newPath, err)
                    }
                    fmt.Printf("Renamed file from %s to %s
", oldPath, newPath)
                }
            }
        }
    }
    return nil
}

func main() {
    // 定义重命名规则
    rules := []RenameRule{
        {SearchPattern: "old", ReplaceWith: "new"}, // 将文件名中的"old"替换为"new"
        {SearchPattern: "tmp", ReplaceWith: "final"}, // 将文件名中的"tmp"替换为"final"
    }

    // 指定目录
    directory := "/path/to/your/directory"

    // 执行文件重命名
    if err := RenameFiles(directory, rules); err != nil {
        fmt.Printf("Error occurred: %s
", err)
    } else {
        fmt.Println("Files renamed successfully.")
    }
}
