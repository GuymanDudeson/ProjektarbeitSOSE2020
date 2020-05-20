package Projektarbeit_Aufgabe4

import "strings"

type Exp interface {
	eval() int
	pretty() string
}

type IntExpInterface interface {
	newIntExp(val int) IntExp
}

type PlusExpInterface interface {
	newPlusExp(e1 Exp, e2 Exp) PlusExp
}

type MultExpInterface interface {
	newMultExp(e1 Exp, e2 Exp) MultExp
}


type IntExp struct {
	val int
}

func (exp IntExp) eval() int{
	return exp.val
}

func (exp IntExp) pretty() string{
	return string(exp.val)
}

type PlusExp struct {
	e1 Exp
	e2 Exp
}

func (exp PlusExp) eval() int{
	return exp.e1.eval() + exp.e2.eval()
}

func (exp PlusExp) pretty() string{
	var s = strings.Builder{}
	s.WriteString("( ")
	s.WriteString(exp.e1.pretty())
	s.WriteString(" + ")
	s.WriteString(exp.e2.pretty())
	s.WriteString(" )")
	return s.String()
}

type MultExp struct {
	e1 Exp
	e2 Exp
}

func (exp MultExp) eval() int{
	return exp.e1.eval() * exp.e2.eval()
}

func (exp MultExp) pretty() string{
	var s = strings.Builder{}
	s.WriteString("( ")
	s.WriteString(exp.e1.pretty())
	s.WriteString(" * ")
	s.WriteString(exp.e2.pretty())
	s.WriteString(" )")
	return s.String()
}