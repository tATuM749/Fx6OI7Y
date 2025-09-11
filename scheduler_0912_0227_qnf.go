// 代码生成时间: 2025-09-12 02:27:30
// scheduler.go

package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler defines a struct to encapsulate the cron schedule
type Scheduler struct {
    Cron *cron.Cron
}

// NewScheduler creates a new instance of Scheduler
func NewScheduler() *Scheduler {
    c := cron.New()
    return &Scheduler{Cron: c}
}

// AddJob adds a new job to the scheduler
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    _, err := s.Cron.AddFunc(spec, cmd)
    if err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.Cron.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.Cron.Stop()
}

// ExampleJob is a sample job that prints the current time
func ExampleJob() {
    fmt.Println("Job executed at", time.Now().Format(time.RFC1123))
}

func main() {
    // Create a new scheduler
    scheduler := NewScheduler()

    // Define a job that will be executed every minute at 30 seconds
    _, err := scheduler.AddJob("@every 1m30s", ExampleJob)
    if err != nil {
        log.Fatal("Failed to add job: ", err)
    }

    // Start the scheduler
    scheduler.Start()

    // Keep the main goroutine running to allow the scheduler to execute jobs
    select {}
}