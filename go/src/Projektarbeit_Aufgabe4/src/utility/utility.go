package utility

import "Projektarbeit_Aufgabe4/src/ast"

type Optional struct {
	b bool
	val ast.Exp
}

func NewEmptyOptional() Optional {
	return Optional{
		b:   false,
		val: nil,
	}
}

func NewOptional(val ast.Exp) Optional{
	return Optional{
		b:   true,
		val: val,
	}
}

func(o Optional) IsJust() bool{
	return o.b
}

func(o Optional) IsNothing() bool{
	return !o.b
}

func(o Optional) FromJust() ast.Exp{
	return o.val
}



