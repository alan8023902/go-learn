package order

import (
	"fmt"
	"golearn/src/user"
)

type Orders struct {
	ID      string
	Amount  float64
	Paid    bool
	UserRef *user.Users
}

// 支付订单

func (o *Orders) Pay() {
	if o.Paid {
		fmt.Println("订单已经支付，请勿重复付款！")
		return
	}

	if o.Amount > o.UserRef.Balance {
		fmt.Println("余额不足，支付失败！")
		return
	}

	o.UserRef.Balance -= o.Amount
	o.Paid = true
	fmt.Printf("✅ 订单 %s 支付成功，扣款 %.2f 元，剩余余额 %.2f 元\n",
		o.ID, o.Amount, o.UserRef.Balance)
}

// 打印订单详情

func (o Orders) PrintOrder() {
	fmt.Printf("📦 订单号: %s | 金额: %.2f | 支付状态: %v | 用户: %s\n",
		o.ID, o.Amount, o.Paid, o.UserRef.Name)
}

// 创建订单
func NewOrder(id string, amout float64, user *user.Users) Orders {
	return Orders{ID: id, Amount: amout, Paid: false, UserRef: user}
}
