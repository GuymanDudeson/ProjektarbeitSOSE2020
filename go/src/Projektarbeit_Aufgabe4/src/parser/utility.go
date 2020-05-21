package parser

type Optional struct {
	b bool
	val Exp
}

func NewEmptyOptional() Optional {
	return Optional{
		b:   false,
		val: nil,
	}
}

func NewOptional(val Exp) Optional {
	return Optional{
		b:   true,
		val: val,
	}
}

func(o Optional) isJust() bool{
	return o.b
}

func(o Optional) isNothing() bool{
	return !o.b
}

func(o Optional) fromJust() Exp{
	return o.val
}



