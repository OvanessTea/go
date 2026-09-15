package main

import (
	"errors"
	"fmt"
)

func main() {
	divide(10, 2)
}

// Practice 1
// func test() {
//     defer fmt.Println("A")
//     defer fmt.Println("B")

//     fmt.Println("C")

//     defer fmt.Println("D")

//     fmt.Println("E")
// }
// Порядок вызова: C -> E -> D -> B -> A

// Practice 2
// func test() {
//     x := 10

//     defer fmt.Println("defer:", x)

//     x = 20

//     fmt.Println("current:", x)
// }
// current: 20
// defer: 10
// Я думаю так, потому что в момент чтения defer строки x был равен 10
// и это запомнилось в памяти

// Practice 3
// func test() {
//     x := 10

//     defer func() {
//         fmt.Println("defer:", x)
//     }()

//     x = 20

//     fmt.Println("current:", x)
// }
// Будет 20. Потому что в момент выхода из функции test() вызывается анонимная функция,
// которая в момент объявления не запоминала состояние, а берёт её в момент вызова
// когда переменная x уже будет равна 20

// Practice 4
// func test() int {
//     x := 10

//     defer func() {
//         x = 20
//     }()

//     return x
// }
// Будет 10. Потому что мы возвращаем переменную до вызова defer функции

// Practice 5
func divide(a, b int) (result int, err error) {
	if b == 0 {
		result = 0
		err = errors.New("division by zero")
	} else {
		result = a / b
		err = nil
	}
	defer fmt.Println("result:", result)
	return
}
