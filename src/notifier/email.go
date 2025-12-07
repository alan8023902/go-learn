package notifier

import "fmt"

/*
  通知接口
*/

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("邮件通知：%s\n", message)
}

func (e EmailNotifier) GetType() string {
	return "Email"
}
