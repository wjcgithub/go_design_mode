package main

// BaseOperator is the base operator
type BaseOperator struct {
	OperatorA, OperatorB int
}

// SetA set operator A
func (o *BaseOperator) SetA(a int) {
	o.OperatorA = a
}

// SetB set operator B
func (o *BaseOperator) SetB(b int) {
	o.OperatorB = b
}

// Operator is the interface of operator
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory is the interface of operator factory
type OperatorFactory interface {
	Create() Operator
}

type AddOperatorFactory struct {
}

func (f *AddOperatorFactory) Create() Operator {
	return &AddOperator{}
}

type AddOperator struct {
	*BaseOperator
}

func (o *AddOperator) Result() int {
	return o.OperatorA + o.OperatorB
}

type SubOperatorFactory struct {
}

func (f *SubOperatorFactory) Create() Operator {
	return &SubOperator{}
}

type SubOperator struct {
	*BaseOperator
}

func (o *SubOperator) Result() int {
	return o.OperatorA - o.OperatorB
}

func main() {
	var factory OperatorFactory
	factory = &AddOperatorFactory{}
	operator := factory.Create()
	operator.SetA(1)
	operator.SetB(2)
	println(operator.Result())

	factory = &SubOperatorFactory{}
	operator = factory.Create()
	operator.SetA(1)
	operator.SetB(2)
	println(operator.Result())
}

// Output:
// 3
// -1

// 对上面的结构给我画一下UML类图
