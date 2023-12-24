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

// Factory method
func CreateNotifier(t string) Notifier {
	if t == "sms" {
		return SmsNotifier{}
	}
	return EmailNotifier{}
}

func main() {
	service := Service{
		notifier: CreateNotifier("sms"),
	}

	service.SendNotification()
}
