package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Email string
}
type SMSNotifier struct {
	Phone string
}

type NotificationService struct {
	Notifier
}

func main() {
	emailService := NotificationService{
		Notifier: EmailNotifier{
			Email: "alice@example.com",
		},
	}

	smsService := NotificationService{
		Notifier: SMSNotifier{
			Phone: "+123456789",
		},
	}

	emailService.Send("Hello")
	smsService.Send("Hello")
}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email to alice@example.com:", message)
}
func (s SMSNotifier) Notify(message string) {
	fmt.Println("Sending SMS to +123456789:", message)
}

func (s NotificationService) Send(message string) {
	s.Notifier.Notify(message)
}

// ┌──────────────┐
// │  Notifier    │
// │              │
// │ Notify(...)  │
// └───┬──────────┘
// 	   │
// interface contract
// 	   │
// ┌───┴─────────────────────────┐
// ↓                             ↓
// ┌──────────────────┐   ┌──────────────────┐
// │ EmailNotifier    │   │ SMSNotifier      │
// │                  │   │                  │
// │ Notify()         │   │ Notify()         │
// └──────────────────┘   └──────────────────┘
