package user

import "fmt"

type Users struct {
	Name    string
	Email   string
	Balance float64
}

func (u Users) PrintUser() {
	fmt.Printf("👤 用户: %s | 邮箱: %s | 余额: %.2f 元\n", u.Name, u.Email, u.Balance)
}

func (u *Users) ChangeAmount(amout float64) {
	u.Balance += amout
	fmt.Printf("💰 %s 成功充值 %.2f 元，当前余额 %.2f 元\n", u.Name, amout, u.Balance)
}
