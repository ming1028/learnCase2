package main

import "fmt"

func main() {
	wxPay := &WxPay{}
	px := NewPayCtx(wxPay)
	px.Pay()

	thPay := &ThirdPay{}
	px.setPayBehavior(thPay)
	px.Pay()
}

type PayCtx struct {
	payBehavior PayBehavior
	payParams   map[string]interface{}
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

type PayBehavior interface {
	OrderPay(px *PayCtx)
}

type WxPay struct{}

func (*WxPay) OrderPay(px *PayCtx) {
	fmt.Printf("Wx支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用Wx支付进行支付")
}

type ThirdPay struct{}

func (*ThirdPay) OrderPay(px *PayCtx) {
	fmt.Printf("三方支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用三方支付进行支付")
}

func NewPayCtx(p PayBehavior) *PayCtx {
	params := map[string]interface{}{
		"appId": "123213",
		"mchId": 123456,
	}
	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}
