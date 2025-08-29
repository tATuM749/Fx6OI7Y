// 代码生成时间: 2025-08-29 16:10:24
// log_parser.go
package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LogEntry represents a single entry in the log file.
type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

// parseLogLine parses a single line from the log and returns a LogEntry.
func parseLogLine(line string) (LogEntry, error) {
	// Assuming log format: [timestamp] [level] message\
	parts := strings.SplitN(line, " ", 3)
	if len(parts) < 3 {
		return LogEntry{}, fmt.Errorf("invalid log format: %s", line)
	}
	return LogEntry{
		Timestamp: parts[0] + " " + parts[1],
		Level:     parts[2][1:2],
		Message:   strings.Join(parts[2:], " "),
	}, nil
}

// parseLogFile reads a log file and parses each line into a LogEntry.
func parseLogFile(filePath string) ([]LogEntry, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s", err)
	}
	lines := strings.Split(string(content), "\
")
	entries := make([]LogEntry, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}
		entry, err := parseLogLine(line)
		if err != nil {
			log.Printf("failed to parse line: %s, error: %v", line, err)
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

// walkDir recursively walks through directories and finds log files.
func walkDir(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".log") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %s", err)
	}
	return files, nil
}

func main() {
	root := "./logs" // Assuming log files are in the './logs' directory
	files, err := walkDir(root)
	if err != nil {
		log.Fatalf("failed to find log files: %v", err)
	}
	for _, file := range files {
		entries, err := parseLogFile(file)
		if err != nil {
			log.Printf("failed to parse log file: %s, error: %v", file, err)
			continue
		}
		for _, entry := range entries {
			fmt.Printf("%+v\
", entry)
		}
	}
}
