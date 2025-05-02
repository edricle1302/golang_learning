package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	// main.go
	fmt.Println("*** Example Decorator ***")

	// Bắt đầu với Email
	var notifier Notifier = &EmailNotifier{}

	// Bọc thêm Facebook
	notifier = &FacebookDecorator{
		BaseNotifierDecorator{wrapped: notifier},
	}

	// Bọc thêm SMS
	notifier = &SMSDecorator{
		BaseNotifierDecorator{wrapped: notifier},
	}

	// Gửi thông báo với các decorator
	notifier.Send("House is on fire!")

	fmt.Print("*** End of Decorator ***\n\n\n")
}
