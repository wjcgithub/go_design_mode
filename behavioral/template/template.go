package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BankBusinessHandler interface {
	// 排队拿号
	TakeRowNumber()
	// 等位
	WaitInHead()
	// 处理具体业务
	HandleBusiness()
	// 对服务作出评价
	Commentate()
	// 钩子方法，用于在流程里判断是不是VIP， 实现类似VIP不用等的需求
	CheckVipIdentity() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

func (b *BankBusinessExecutor) ExecuteBankBusiness() {
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

// 通用的方法还可以抽象到BaseBusinessHandler里，组合到具体实现类里，减少重复代码（实现类似子类继承抽象类的效果）
func (*DepositBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dh *DepositBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVip
}

func (*DepositBusinessHandler) Commentate() {

	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

// 通用的方法还可以抽象到BaseBusinessHandler里，组合到具体实现类里，减少重复代码（实现类似子类继承抽象类的效果）
type DefaultBusinessHandler struct {
}

func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dbh *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultBusinessHandler) Commentate() {

	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

func main() {
	// 模拟一个存款业务
	depositBusinessHandler := &DepositBusinessHandler{
		DefaultBusinessHandler: &DefaultBusinessHandler{},
		userVip:                true,
	}
	bankBusinessExecutor := NewBankBusinessExecutor(depositBusinessHandler) //解决这个问题
	bankBusinessExecutor.ExecuteBankBusiness()
}
