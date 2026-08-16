package main

import (
	"errors"
	"fmt"
)

func main() {
	statuses := []string{
		"pending",
		"in_progress",
		"completed",
		"failed",
		"cancelled",
		"not_allowed",
	}
	for _, status := range statuses {
		message, err := statusMessage(status)

		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Println(message)
	}
}

func statusMessage(str string) (string, error) {
	switch str {
	case "pending":
		return "Task is waiting", nil
	case "in_progress":
		return "Task is being processed", nil
	case "completed":
		return "Task completed", nil
	case "failed":
		return "Task failed", nil
	case "cancelled":
		return "Task was cancelled", nil
	default:
		return "", errors.New("Unknown task status")
	}
}
