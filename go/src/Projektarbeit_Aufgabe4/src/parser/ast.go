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
	eval() int
	getKind() expType
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

func (exp IntExp) getKind() expType{
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

func (exp PlusExp) eval() int{
	return exp.e1.eval() + exp.e2.eval()
}

func (exp PlusExp) getKind() expType{
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

func (exp MultExp) eval() int{
	return exp.e1.eval() * exp.e2.eval()
}

func (exp MultExp) getKind() expType{
	return MULTEXP
}

func (exp MultExp) pretty() string{
	var s = strings.Builder{}

	if exp.e1.getKind() == 1{
		s.WriteString("(")
		s.WriteString(exp.e1.pretty())
		s.WriteString(")")
	} else {
		s.WriteString(exp.e1.pretty())
	}

	s.WriteString(" * ")

	if exp.e2.getKind() == 1{
		s.WriteString("(")
		s.WriteString(exp.e2.pretty())
		s.WriteString(")")
	} else {
		s.WriteString(exp.e2.pretty())
	}
	return s.String()
}