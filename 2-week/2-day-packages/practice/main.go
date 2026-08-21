package main

func main() {

}

// Task 1
// entities | pkgs
// User		|	user/ -> all about user (including email)
// Product	|	product/ -> all about products (info, stats, etc.)
// Order	|	order/ -> all about ordering (methods to order goods)
// Payment	|	payment/ -> all for paying for goods
// Email	| 	----

// Task 2
// utils/
//     user.go
//     email.go
//     payment.go
//     date.go
//     json.go
// This could provide problem when we lost undestanding what our code does
// utils.DoSomething()
// utils.ValidatePassword()
// utils.ChangeEmail()
// That methods should be separated for different pkgs
// 'cause they have diff entities and should have sep responsability

// Task 3
// user → notification - User has dependencies from Notification
//	user → notification
// notification → user
// import cycle is forbidden. Dependencies have to be linear

// Task 4
// Code has excessive dependence
// We sohuld use interface for decreasing connectedness of pkgs
// If we add interface, code won't think what send notifications.
// It would worry only about Notify
// package user

// type Notifier interface {
//     Notify(message string)
// }

// type Service struct {
//     notifier Notifier
// }

//          user
//           │
//           │ Notifier
//           ▼
//     ┌──────────────┐
//     │              │
// EmailNotifier   SMSNotifier
