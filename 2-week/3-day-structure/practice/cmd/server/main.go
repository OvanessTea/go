package main

import (
	"fmt"

	"example.com/myapp/internal/notification"
	"example.com/myapp/internal/user"
)

func main() {
	u := user.NewUser("Alex", 35)

	emailNotifier := notification.EmailNotifier{}
	userService := user.NewService(emailNotifier)

	userService.SendNotification("Hello")
	fmt.Println(u)
}
