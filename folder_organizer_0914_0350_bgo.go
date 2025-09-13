// 代码生成时间: 2025-09-14 03:50:53
This program will scan a directory and its subdirectories, organizing files into
subfolders based on a predefined set of rules.

@author: Your Name
@date: 2023-04-01
*/

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "log"
)

// FolderOrganizer defines the structure for organizing folders.
type FolderOrganizer struct {
    RootPath string
}

// NewFolderOrganizer creates a new FolderOrganizer instance.
func NewFolderOrganizer(rootPath string) *FolderOrganizer {
    return &FolderOrganizer{
        RootPath: rootPath,
    }
}

// Organize organizes the folder structure based on the rules.
func (f *FolderOrganizer) Organize() error {
    // List all files and directories in the root path.
    files, err := os.ReadDir(f.RootPath)
    if err != nil {
        return fmt.Errorf("failed to read the directory: %w", err)
    }

    for _, file := range files {
        filePath := filepath.Join(f.RootPath, file.Name())
        if file.IsDir() {
            // If it's a directory, recursively call Organize.
            if err := f.Organize(); err != nil {
                return err
            }
        } else {
            // If it's a file, organize it based on its extension.
            if err := f.organizeFile(filePath); err != nil {
                return err
            }
        }
    }
    return nil
}

// organizeFile organizes a single file based on its extension.
func (f *FolderOrganizer) organizeFile(filePath string) error {
    extension := strings.TrimPrefix(filepath.Ext(filePath), ".")
    if extension == "" {
        return nil // Skip files without an extension.
    }

    // Create a directory for the file's extension.
    dirPath := filepath.Join(f.RootPath, strings.ToLower(extension))
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        if err := os.MkdirAll(dirPath, 0755); err != nil {
            return fmt.Errorf("failed to create directory: %w", err)
        }
    }

    // Move the file to the new directory.
    destPath := filepath.Join(dirPath, filepath.Base(filePath))
    if err := os.Rename(filePath, destPath); err != nil {
        return fmt.Errorf("failed to move file: %w", err)
    }
    return nil
}

func main() {
    rootPath := "./" // Set the root path to the current directory.
    organizer := NewFolderOrganizer(rootPath)
    if err := organizer.Organize(); err != nil {
        log.Fatalf("error organizing folders: %s", err)
    }
    fmt.Println("Folder structure organized successfully.")
}