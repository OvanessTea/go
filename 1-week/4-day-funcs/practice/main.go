package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(multiply(5, 4))
	fmt.Println(minMax([]int{5, 2, 9, 1, 7}))
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
	fmt.Println(statusMessage("completed"))

	message, err := statusMessageError("failed")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

}

// Task 1
func multiply(a, b int) int {
	return a * b
}

// Task 2
func minMax(numbers []int) (int, int) {
	var max int
	min := 999
	for _, value := range numbers {
		if value > max {
			max = value
		} else if value < min {
			min = value
		}
	}
	return min, max
}

// Task 3
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

// Task 4
func statusMessage(str string) string {
	switch str {
	case "pending":
		return "Task is waiting"
	case "in_progress":
		return "Task is being processed"
	case "completed":
		return "Task completed"
	case "failed":
		return "Task failed"
	case "cancelled":
		return "Task was cancelled"
	default:
		return "Unknown task status"
	}
}

// Task 5
func statusMessageError(str string) (string, error) {
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

// To think
// value == 10, err == nil, нет, fmt.Printlin(value) -> 10
// func test() (int, error) {
//     return 10, nil
// }

// func main() {
//     value, err := test()

//     if err != nil {
//         return
//     }

//     fmt.Println(value)
// }
