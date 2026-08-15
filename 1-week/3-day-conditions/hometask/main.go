package main

import "fmt"

func main() {
	statuses := []string{
		"pending",
		"in_progress",
		"completed",
		"failed",
		"cancelled",
	}
	for _, status := range statuses {
		statusCheck(status)
	}
}

func statusCheck(str string) {
	switch str {
	case "pending":
		fmt.Println("Task is waiting")
	case "in_progress":
		fmt.Println("Task is being processed")
	case "completed":
		fmt.Println("Task completed")
	case "failed":
		fmt.Println("Task failed")
	case "cancelled":
		fmt.Println("Task was cancelled")
	default:
		fmt.Println("Unknown task status")
	}
}
