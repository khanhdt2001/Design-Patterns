package main

import "fmt"

type Notifier interface {
	Notify()
}

type EmailNotifier struct {
}

func (EmailNotifier) Notify() {
	fmt.Println("Email notify")
}

type SmsNotifier struct {
}

func (SmsNotifier) Notify() {
	fmt.Println("SMS notify")
}

type Service struct {
	notifier Notifier
}

func (s Service) SendNotification() {
	s.notifier.Notify()
}

func main() {
	service := Service{
		notifier: EmailNotifier{},
	}

	service.SendNotification()
}
