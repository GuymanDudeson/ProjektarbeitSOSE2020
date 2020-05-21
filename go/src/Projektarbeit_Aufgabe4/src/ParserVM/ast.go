package ParserVM

import "strings"

type Exp interface {
	Eval() int
	Pretty() string
}

type IntExp struct {
	val int
}

func NewIntExp(val int) IntExp {
	return IntExp{val: val}
}

func (exp IntExp) Eval() int{
	return exp.val
}

func (exp IntExp) Pretty() string{
	return string(exp.val)
}

type PlusExp struct {
	e1 Exp
	e2 Exp
}

func NewPlusExp(e1 Exp, e2 Exp) PlusExp {
	return PlusExp{
		e1: e1,
		e2: e2,
	}
}

func (exp PlusExp) Eval() int{
	return exp.e1.Eval() + exp.e2.Eval()
}

func (exp PlusExp) Pretty() string{
	var s = strings.Builder{}
	s.WriteString("( ")
	s.WriteString(exp.e1.Pretty())
	s.WriteString(" + ")
	s.WriteString(exp.e2.Pretty())
	s.WriteString(" )")
	return s.String()
}

type MultExp struct {
	e1 Exp
	e2 Exp
}

func NewMultExp(e1 Exp, e2 Exp) MultExp {
	return MultExp{
		e1: e1,
		e2: e2,
	}
}

func (exp MultExp) Eval() int{
	return exp.e1.Eval() * exp.e2.Eval()
}

func (exp MultExp) Pretty() string{
	var s = strings.Builder{}
	s.WriteString("( ")
	s.WriteString(exp.e1.Pretty())
	s.WriteString(" * ")
	s.WriteString(exp.e2.Pretty())
	s.WriteString(" )")
	return s.String()
}