package notification

import "fmt"

type EmailNotifier struct{}

func (EmailNotifier) Notify(message string) {
	fmt.Println("email:", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Notify(message string) {
	fmt.Println("sms:", message)
}
