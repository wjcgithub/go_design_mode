package main

import "fmt"

// PayCtx is the context of pay
type PayCtx struct {
	payBehavior PayBehavior
	payParams   map[string]interface{}
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) pay() {
	px.payBehavior.Pay(px)
}

// NewPayCtx is the constructor of PayCtx
func NewPayCtx(p PayBehavior) *PayCtx {
	params := map[string]interface{}{
		"uid":    123,
		"amount": 100,
	}
	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}

// PayBehavior is the interface of pay behavior
type PayBehavior interface {
	Pay(px *PayCtx)
}

// PayByWechat is the pay behavior of ali
type PayByAli struct {
}

func (p *PayByAli) Pay(px *PayCtx) {
	fmt.Println("Pay by Ali Params", px.payParams)
	fmt.Println("Pay by Ali")
}

// PayByWechat is the pay behavior of wechat
type PayByWechat struct {
}

func (p *PayByWechat) Pay(px *PayCtx) {
	fmt.Println("Pay by Wechat Params", px.payParams)
	fmt.Println("Pay by Wechat")
}

func main() {
	px := NewPayCtx(&PayByAli{})
	px.pay()

	px.setPayBehavior(&PayByWechat{})
	px.pay()
}

// Output:
// Pay by Ali Params map[amount:100 uid:123]
// Pay by Ali
// Pay by Wechat Params map[amount:100 uid:123]
// Pay by Wechat
