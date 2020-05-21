package parser

import (
	"strconv"
	"strings"
)

type Exp interface {
	eval() int
	pretty() string
}

type IntExp struct {
	val int
}

func NewIntExp(val int) IntExp {
	return IntExp{val: val}
}

func (exp IntExp) eval() int{
	return exp.val
}

func (exp IntExp) pretty() string{
	return strconv.Itoa(exp.val)
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

func (exp PlusExp) eval() int{
	return exp.e1.eval() + exp.e2.eval()
}

func (exp PlusExp) pretty() string{
	var s = strings.Builder{}
	s.WriteString("(")
	s.WriteString(exp.e1.pretty())
	s.WriteString(" + ")
	s.WriteString(exp.e2.pretty())
	s.WriteString(")")
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

func (exp MultExp) eval() int{
	return exp.e1.eval() * exp.e2.eval()
}

func (exp MultExp) pretty() string{
	var s = strings.Builder{}
	s.WriteString("(")
	s.WriteString(exp.e1.pretty())
	s.WriteString(" * ")
	s.WriteString(exp.e2.pretty())
	s.WriteString(")")
	return s.String()
}