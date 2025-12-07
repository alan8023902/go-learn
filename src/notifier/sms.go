package notifier

import "fmt"

type SmsNotifier struct{}

func (s SmsNotifier) Send(message string) {
	fmt.Printf("短信通知：%s\n", message)
}

func (s SmsNotifier) GetType() string {
	return "SMS"
}
