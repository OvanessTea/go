package main

import "fmt"

func main() {
	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	// while
	counter := 0

	for counter < 5 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("-----------")

	// while true + break
	flag := false
	for {
		if flag {
			break
		}
		fmt.Println(flag)
		flag = true
	}

	fmt.Println("-----------")

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("-----------")

	// for range
	names := []string{"Alice", "Bob", "Alex"}
	for index, name := range names {
		fmt.Println(index, name)
	}
	for _, name := range names { // _ - blank identifier
		fmt.Println(name)
	}

	fmt.Println("-----------")

	// for map range
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	for name, age := range ages { // !!no guarantee the same order!!
		fmt.Println(name, age)
	}

	fmt.Println("-----------")

	// switch
	status := "pending"
	switch status {
	case "pending":
		fmt.Println("Waiting") // no need to use break in each case
		// fallthrough - to fall through cases. In most cases useless.
	case "active":
		fmt.Println("Running")
	case "completed":
		fmt.Println("Done")
	default:
		fmt.Println("Unknown")
	}

	fmt.Println("-----------")
}
