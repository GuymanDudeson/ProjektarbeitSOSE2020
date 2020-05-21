package parser

import (
	"strconv"
	"strings"
)

type expType int
const (
	INTEXP expType = 0
	PLUSEXP expType = 1
	MULTEXP expType = 2
)


type Exp interface {
	GetE1() Exp
	GetE2() Exp
	Eval() int
	GetKind() expType
	pretty() string
}

type IntExp struct {
	val int
}

func NewIntExp(val int) IntExp {
	return IntExp{val: val}
}

func (exp IntExp) GetE1() Exp{
	return exp
}

func (exp IntExp) GetE2() Exp{
	return exp
}

func (exp IntExp) Eval() int{
	return exp.val
}

func (exp IntExp) GetKind() expType{
	return INTEXP
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

func(exp PlusExp) GetE1() Exp{
	return exp.e1
}

func(exp PlusExp) GetE2() Exp{
	return exp.e2
}

func (exp PlusExp) Eval() int{
	return exp.e1.Eval() + exp.e2.Eval()
}

func (exp PlusExp) GetKind() expType{
	return PLUSEXP
}

func (exp PlusExp) pretty() string{
	var s = strings.Builder{}
	s.WriteString(exp.e1.pretty())
	s.WriteString(" + ")
	s.WriteString(exp.e2.pretty())
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

func(exp MultExp) GetE1() Exp{
	return exp.e1
}

func(exp MultExp) GetE2() Exp{
	return exp.e2
}

func (exp MultExp) Eval() int{
	return exp.e1.Eval() * exp.e2.Eval()
}

func (exp MultExp) GetKind() expType{
	return MULTEXP
}

func (exp MultExp) pretty() string{
	var s = strings.Builder{}

	if exp.e1.GetKind() == 1{
		s.WriteString("(")
		s.WriteString(exp.e1.pretty())
		s.WriteString(")")
	} else {
		s.WriteString(exp.e1.pretty())
	}

	s.WriteString(" * ")

	if exp.e2.GetKind() == 1{
		s.WriteString("(")
		s.WriteString(exp.e2.pretty())
		s.WriteString(")")
	} else {
		s.WriteString(exp.e2.pretty())
	}
	return s.String()
}