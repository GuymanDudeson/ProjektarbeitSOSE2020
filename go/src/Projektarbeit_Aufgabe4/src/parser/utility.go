package parser

type Optional struct {
	b   bool
	Val Exp
}

func NewEmptyOptional() Optional {
	return Optional{
		b:   false,
		Val: nil,
	}
}

func NewOptional(val Exp) Optional {
	return Optional{
		b:   true,
		Val: val,
	}
}

func(o Optional) isJust() bool{
	return o.b
}

func(o Optional) isNothing() bool{
	return !o.b
}

func(o Optional) fromJust() Exp{
	return o.Val
}



