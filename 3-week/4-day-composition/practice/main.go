package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Address struct {
	City   string
	Street string
}

type User2 struct {
	Name string
	Address
}

func main() {
	user, err := NewUser("Alex", 35)
	if err != nil {
		fmt.Println(err)
		// return - I commented returns for next tasks
	}

	fmt.Println(user)

	user, err = NewUser("", 35)
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Println(user)

	user, err = NewUser("Alex", -5)
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Println(user)

	user2 := User2{
		Name: "Alex",
		Address: Address{
			City:   "Berlin",
			Street: "Main Street",
		},
	}

	fmt.Println(user2.Name)
	fmt.Println(user2.City)
	fmt.Println(user2.Street)

	fmt.Println(user2.FullAddress())
}

func NewUser(name string, age int) (User, error) {
	if name == "" {
		return User{}, fmt.Errorf("name is required")
	}

	if age < 0 {
		return User{}, fmt.Errorf("invalid age: %d", age)
	}

	return User{Name: name, Age: age}, nil
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s", a.City, a.Street)
}

// Task 4 to think
// Потому что происходит композиция метода от Address{}. Поэтому у User{} есть метод Validate()

// Самопроверка:
// 1. Нет. Это обычная функция
// 2. Потому что конструктор помогает добавлять валидацию при создании структуры
// 3. Когда одна структура является параметром другой
// 4. Запись Address -> Address. Это позволяет использовать promoted поля и методы
// 5. Параметры "дочерней" структры по композиции передаются "родительскому".
// Что позволяет вызывать: user.City вместо user.Address.City
// 6. Методы, так же как и поля могут быть использованы не только дочерней структурой,
// но и родительской
// 7. Нет
// 8. Потому что происходит композиция метода. То есть: promoted method
// 9. Думаю, такие поля/методы просто не могут быть promoted. у каждой сущности эти поля/методы буду свои.
