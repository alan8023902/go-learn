package notifier

/*
  通知接口
*/

type Notifier interface {
	Send(message string)
	GetType() string
}
