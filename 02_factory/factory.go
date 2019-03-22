package factory

//工厂模式，抽象

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

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

//加法版，延迟实现
type OperatorPlus struct {
	*OperatorBase
}

//这里不是指针接收者,SetA 方法是OperatorPlus调用的
func (op OperatorPlus) Result() int {
	return op.a + op.b
}

type PlusOperatorFactory struct{}

//工厂方法
func (p *PlusOperatorFactory) Create() Operator {
	return OperatorPlus{
		OperatorBase: &OperatorBase{},
	}
}

//减法版
type OperatorMinus struct {
	*OperatorBase
}

func (op OperatorMinus) Result() int {
	return op.a - op.b
}

type MinusOperatorFactory struct{}

func (mop *MinusOperatorFactory) Create() Operator {
	return OperatorMinus{
		OperatorBase: &OperatorBase{},
	}
}
