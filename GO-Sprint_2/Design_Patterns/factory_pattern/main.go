package main

import "fmt"

type Notification interface {
	SendNotification(message string) string
}

type SMSNotification struct {
}

func (s *SMSNotification) SendNotification(message string) string {
	return fmt.Sprintf("Send the sms : %v", message)
}

type EmailNotification struct {
}

func (e *EmailNotification) SendNotification(message string) string {
	return fmt.Sprintf("Send the email : %v", message)
}

type PushNotification struct {
}

func (p *PushNotification) SendNotification(message string) string {
	return fmt.Sprintf("Send the push : %v", message)
}

func NewNotificationFactory(notificationType string) (Notification, error) {
	switch notificationType {
	case "sms":
		return &SMSNotification{}, nil
	case "email":
		return &EmailNotification{}, nil
	case "push":
		return &PushNotification{}, nil
	default:
		return nil, fmt.Errorf("the unknown notification type :%s", notificationType)
	}
}

func main() {
	notification, err := NewNotificationFactory("sms")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(notification.SendNotification("Hello via sms!"))

	notification, err = NewNotificationFactory("email")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(notification.SendNotification("Hello via email!"))

	notification, err = NewNotificationFactory("push")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(notification.SendNotification("Hello via push!"))
}
