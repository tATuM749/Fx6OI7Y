// 代码生成时间: 2025-09-03 18:28:08
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// BackupAndSync 是一个结构体，用于配置文件备份和同步工具
type BackupAndSync struct {
    SourceDir  string
    DestinationDir string
    // 可以添加更多的字段以扩展功能
}

// NewBackupAndSync 创建一个新的 BackupAndSync 实例
func NewBackupAndSync(sourceDir, destinationDir string) *BackupAndSync {
    return &BackupAndSync{
        SourceDir:  sourceDir,
        DestinationDir: destinationDir,
    }
}

// SyncFiles 同步源目录下的文件到目标目录
func (b *BackupAndSync) SyncFiles() error {
    // 获取源目录的文件列表
    files, err := ioutil.ReadDir(b.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        srcPath := filepath.Join(b.SourceDir, file.Name())
        destPath := filepath.Join(b.DestinationDir, file.Name())

        // 检查目标目录中是否存在文件，如果存在则跳过
        if _, err := os.Stat(destPath); os.IsNotExist(err) {
            // 复制文件
            if err := b.copyFile(srcPath, destPath); err != nil {
                return fmt.Errorf("failed to copy file %s: %w", file.Name(), err)
            }
        } else {
            fmt.Printf("File %s already exists in destination. Skipping.
", file.Name())
        }
    }
    return nil
}

// BackupFiles 备份源目录下的文件到目标目录
func (b *BackupAndSync) BackupFiles() error {
    // 获取源目录的文件列表
    files, err := ioutil.ReadDir(b.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        srcPath := filepath.Join(b.SourceDir, file.Name())
        destPath := filepath.Join(b.DestinationDir, file.Name()) + ".bak"

        // 复制文件
        if err := b.copyFile(srcPath, destPath); err != nil {
            return fmt.Errorf("failed to backup file %s: %w", file.Name(), err)
        }
    }
    return nil
}

// copyFile 复制单个文件
func (b *BackupAndSync) copyFile(src, dest string) error {
    // 打开源文件
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer srcFile.Close()

    // 创建目标文件
    destFile, err := os.Create(dest)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer destFile.Close()

    // 复制文件内容
    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return fmt.Errorf("failed to copy file content: %w", err)
    }
    return nil
}

func main() {
    // 示例用法
    sourceDir := "./source"
    destinationDir := "./destination"
    backupSync := NewBackupAndSync(sourceDir, destinationDir)

    // 同步文件
    if err := backupSync.SyncFiles(); err != nil {
        fmt.Printf("Error syncing files: %v
", err)
    } else {
        fmt.Println("Files synced successfully.
")
    }

    // 备份文件
    if err := backupSync.BackupFiles(); err != nil {
        fmt.Printf("Error backing up files: %v
", err)
    } else {
        fmt.Println("Files backed up successfully.
")
    }
}