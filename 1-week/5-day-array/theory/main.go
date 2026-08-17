package main

import "fmt"

func main() {
	// ARRAY
	// array has a fix length. Forbidden to change its size
	arr := [3]int{3, 5, 7}

	// length
	fmt.Println(len(arr))

	// SLICE
	// slice - is a dynamic size list
	slice := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(slice)

	// Array and Slice have a diff type
	fmt.Printf("%T\n", arr)   //[3]int
	fmt.Printf("%T\n", slice) //[]string

	// Allow yo add new element to the Slice
	// append RETURNS new value. Has to be saved into var
	slice = append(slice, "Max")
	fmt.Println(slice)

	// Cap shows how many els could be added to the slice before getting new backing array
	fmt.Println(cap(slice))

	// backing array
	// 	slice
	//  ├── pointer ──────┐
	//  ├── len = 3       │
	//  └── cap = 5       │
	//                     ↓
	//               backing array
	//               [10 20 30 _ _]

	// several slices could base on the same backing array
	numbers := []int{10, 20, 30, 40}
	part := numbers[1:3] // slice expression
	// numbers[start:end] -> start <= index < end
	// 	index:  0   1   2   3
	// value: 10  20  30  40
	//            └──────┘
	part[0] = 999 // changing one slice from same backing array will change others

	fmt.Println(numbers)

	numbers2 := []int{1, 2, 3, 4}
	part2 := numbers2[:2]
	part2 = append(part2, 100)

	fmt.Println(numbers2)

	// Undependent copy
	source := []int{1, 2, 3, 4}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(source)
	fmt.Println(destination)

	destination[0] = 999
	fmt.Println(source)
	fmt.Println(destination)

	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10) // capacity is optional
	fmt.Println(cap(slice1))     // len == 5, cap == 5
	fmt.Println(cap(slice2))     // len == 5, cap == 10

	// MAP
	ages := map[string]int{
		"Alice": 30,
		"Mike":  28,
		"Bob":   35,
	}
	fmt.Println(ages["Alice"])
	ages["Dave"] = 40  // Allows to add new elem
	ages["Alice"] = 25 // if elem exists -> update it
	fmt.Println(ages)
	delete(ages, "Bob") // Allows to delete elem
	fmt.Println(ages)

	age, ok := ages["Alice"] // ok - Does elem exists
	fmt.Println(age, ok)

	_, ok = ages["Unknown"]
	if !ok {
		fmt.Println("User not found")
	}

	// range
	for name, age := range ages { // Order  not guaranteed
		fmt.Println(name, age)
	}

	// To create an empy map
	map1 := make(map[string]int)
	map2 := map[string]int{}
	fmt.Println(map1, map2)
}
