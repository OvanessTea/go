package main

import "fmt"

func main() {
	// Task 1
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers[2])
	fmt.Println(len(numbers))
	for _, val := range numbers {
		fmt.Println(val)
	}

	// Task 2
	numbers2 := []int{1, 2, 3}
	numbers2 = append(numbers2, 4)
	numbers2 = append(numbers2, 5)
	numbers2 = append(numbers2, 6)

	fmt.Println(len(numbers2))
	fmt.Println(cap(numbers2))
	fmt.Println(numbers2)

	// Task 3
	numbers3 := []int{10, 20, 30, 40, 50}
	part3 := numbers3[1:4]
	// part3 == [20, 30, 40]
	// len == 3
	fmt.Println(part3)
	fmt.Println(len(part3))
	part3[0] = 999
	// part3 == [999, 30, 40]
	// numbers3 == [10, 999, 30, 40, 50]
	fmt.Println(part3)
	fmt.Println(numbers3)

	// Task 4
	source := []int{1, 2, 3}
	source_copy := make([]int, len(source))
	copy(source_copy, source)
	source_copy[0] = 999
	fmt.Println(source)      // [1, 2, 3]
	fmt.Println(source_copy) // [999, 2, 3]

	// Task 5
	users := map[int64]string{
		1: "Alice",
		2: "Bob",
		3: "Charlie",
	}

	user, ok := users[2]
	if ok {
		fmt.Println(user)
	}

	user, ok = users[999]
	if !ok {
		fmt.Println("User not found")
	}

	// Task to think 1
	// numbers == [10, 999, 30, 40]
	// part == [999, 30]

	// Task to think 2
	// numbers == [1, 2, 100]
	// part == [1, 2, 100]

	// Task to think 3
	// a == 0, b == 0
	// a == 0, okA == true, b == 0, okB == false
}
