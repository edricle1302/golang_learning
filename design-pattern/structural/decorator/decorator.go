package decorator

import "fmt"

// Notifier là interface chung
type Notifier interface {
	Send(message string)
}

// EmailNotifier là thành phần gốc
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending EMAIL with message:", message)
}

// Base Decorator chứa tham chiếu đến Notifier khác
type BaseNotifierDecorator struct {
	wrapped Notifier
}

func (b *BaseNotifierDecorator) Send(message string) {
	b.wrapped.Send(message)
}

// SMS Concrete Decorator thêm chức năng gửi SMS
type SMSDecorator struct {
	BaseNotifierDecorator
}

func (s *SMSDecorator) Send(message string) {
	s.BaseNotifierDecorator.Send(message)
	fmt.Println("Sending SMS with message:", message)
}

// Facebook Concrete Decorator thêm chức năng gửi Facebook
type FacebookDecorator struct {
	BaseNotifierDecorator
}

func (f *FacebookDecorator) Send(message string) {
	f.BaseNotifierDecorator.Send(message)
	fmt.Println("Sending Facebook message:", message)
}
