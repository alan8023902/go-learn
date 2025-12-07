package notifier

import "fmt"

type AINotifier struct{}

func (a AINotifier) Send(message string) {
	fmt.Printf("AI助手通知: %s\n", message)
}

func (a AINotifier) GetType() string {
	return "AI"
}
