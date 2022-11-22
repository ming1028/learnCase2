package main

func main() {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		panic("err")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 3, 2) != 1 {
		panic("err")
	}
}

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 工厂接口，返回Operator结构体
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

// 入口方法
func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
