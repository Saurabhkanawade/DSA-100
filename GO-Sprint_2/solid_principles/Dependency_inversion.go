package main

import "fmt"

//with dependency inversion

type MessageSender interface {
	Send(message string, to string)
}

type EmailService struct {
}

func (s *EmailService) Send(message string, to string) {
	fmt.Println("sending the email")
}

type SMSService struct {
}

func (n *SMSService) Send(email string, to string) {
	n.Send(email, to)
}

type NotificationService struct {
	messageSender MessageSender
}

func (s *NotificationService) Notify(msg string, to string) {
	s.messageSender.Send(msg, to)
}

// EmailService without dependency inversion
//type EmailService struct {
//}
//
//func (s *EmailService) Send(message string, to string) {
//	fmt.Println("sending the email")
//}
//
//type NotificationService struct {
//	emailService *EmailService
//}
//
//func (n *NotificationService) Notify(email string) {
//	n.emailService.Send("hi","email.com")
//
//}

func main() {

}
